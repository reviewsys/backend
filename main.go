package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	deliveryGrpc "github.com/reviewsys/backend/app/delivery/grpc"
	appRepo "github.com/reviewsys/backend/app/repository"
	appUcase "github.com/reviewsys/backend/app/usecase"
	"github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
	cfg "github.com/reviewsys/backend/config"
	"google.golang.org/grpc"
)

var (
	config       cfg.Config
	logrusLogger *logrus.Logger
	customFunc   grpc_logrus.CodeToLevel
)

func init() {
	config = cfg.NewViperConfig()

	if config.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	}
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {

	// Logrus entry is used, allowing pre-definition of certain fields by the user.
	logrusLogger = logrus.New()
	logrusLogger.SetFormatter(&logrus.JSONFormatter{})
	logrusEntry := logrus.NewEntry(logrusLogger)
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
	grpc_logrus.ReplaceGrpcLogger(logrusEntry)

	dbHost := config.GetString(`database.host`)
	dbPort := config.GetString(`database.port`)
	dbUser := config.GetString(`database.user`)
	dbPass := config.GetString(`database.pass`)
	dbName := config.GetString(`database.name`)
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Tokyo")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := sql.Open(`postgres`, dsn)
	if err != nil && config.GetBool("debug") {
		fmt.Println(err)
	}
	defer dbConn.Close()

	ar := appRepo.NewDatabaseUserRepository(dbConn)
	au := appUcase.NewUserUsecase(ar)
	list, err := net.Listen("tcp", config.GetString("server.address"))
	if err != nil {
		fmt.Println("SOMETHING HAPPEN")
	}

	s := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_prometheus.StreamServerInterceptor,
			grpc_logrus.StreamServerInterceptor(logrusEntry, opts...),
			grpc_recovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_prometheus.UnaryServerInterceptor,
			grpc_logrus.UnaryServerInterceptor(logrusEntry, opts...),
			grpc_recovery.UnaryServerInterceptor(),
		)),
	)

	deliveryGrpc.NewAppServerGrpc(s, au)
	fmt.Println("Server Run at ", config.GetString("server.address"))

	grpc_prometheus.Register(s)
	// Register Prometheus metrics handler.
	http.Handle("/metrics", promhttp.Handler())

	err = s.Serve(list)
	if err != nil {
		fmt.Println("Unexpected Error", err)
	}
}
