package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	deliveryGrpc "github.com/reviewsys/backend/app/delivery/grpc"
	appRepo "github.com/reviewsys/backend/app/repository"
	appUcase "github.com/reviewsys/backend/app/usecase"

	cfg "github.com/reviewsys/backend/config"
	"google.golang.org/grpc"
)

var config cfg.Config

func init() {
	config = cfg.NewViperConfig()

	if config.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	}
	log.SetFlags(log.LstdFlags | log.Lshortfile)

}

func main() {
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
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
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
