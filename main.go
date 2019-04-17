package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"

	"github.com/reviewsys/backend/app/interface/rpc"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/reviewsys/backend/app/registry"
	log "github.com/sirupsen/logrus"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	cfg "github.com/reviewsys/backend/config"
	"google.golang.org/grpc"
)

var (
	config     cfg.Config
	customFunc grpc_logrus.CodeToLevel
)

func init() {
	config = cfg.NewViperConfig()

	if config.GetBool(`debug`) {
		log.SetLevel(log.DebugLevel)
		log.Debug("Service RUN on DEBUG mode")
	} else {
		log.SetLevel(log.InfoLevel)
	}
}

func main() {
	// Logrus entry is used, allowing pre-definition of certain fields by the user.
	logger := log.NewEntry(log.New())
	log.SetFormatter(&log.JSONFormatter{})
	// Shared options for the logger, with a custom gRPC code to log level function.

	customFunc = grpc_logrus.DefaultCodeToLevel
	opts := []grpc_logrus.Option{
		grpc_logrus.WithLevels(customFunc),
	}

	// Shared options for the logger, with a custom duration to log field function.
	//opts := []grpc_logrus.Option{
	//	grpc_logrus.WithDurationField(func(duration time.Duration) (key string, value interface{}) {
	//		return "grpc.time_ns", duration.Nanoseconds()
	//	}),
	//}

	// Make sure that log statements internal to gRPC library are logged using the logrus Logger as well.
	grpc_logrus.ReplaceGrpcLogger(logger)

	dbHost := config.GetString(`database.host`)
	dbPort := config.GetString(`database.port`)
	dbUser := config.GetString(`database.user`)
	dbPass := config.GetString(`database.pass`)
	dbName := config.GetString(`database.name`)

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		dbHost,
		dbPort,
		dbUser,
		dbName,
		dbPass,
	)

	ctn, err := registry.NewContainer(dsn)
	if err != nil {
		log.Errorf("failed to build container: %v", err)
	}
	list, err := net.Listen("tcp", config.GetString("server.address"))
	if err != nil {
		log.Errorf("SOMETHING HAPPEN: %v", err)
	}

	server := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_prometheus.StreamServerInterceptor,
			grpc_logrus.StreamServerInterceptor(logger, opts...),
			grpc_recovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_prometheus.UnaryServerInterceptor,
			grpc_logrus.UnaryServerInterceptor(logger, opts...),
			grpc_recovery.UnaryServerInterceptor(),
		)),
	)

	rpc.Apply(server, ctn)

	grpc_prometheus.Register(server)
	// Register Prometheus metrics handler.
	http.Handle("/metrics", promhttp.Handler())

	go func() {
		log.Info("Server Run at ", config.GetString("server.address"))
		server.Serve(list)
		if err != nil {
			log.Errorf("Unexpected Error: %v", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping grpc server...")
	server.GracefulStop()
	ctn.Clean()
}
