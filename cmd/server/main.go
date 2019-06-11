package main

import (
	"context"
	"fmt"
	"net"
	"os"

	cfg "github.com/reviewsys/backend/config"

	"github.com/infobloxopen/atlas-app-toolkit/gorm/resource"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	pb "github.com/reviewsys/backend/app/interface/rpc/v1.0/protocol"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var (
	config cfg.Config
)

type Server struct {
	db *gorm.DB
}

func NewServer(db *gorm.DB) *Server {
	return &Server{
		db: db,
	}
}

func (s *Server) Create(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	logrus.Debugf("req: %v", req)
	user, err := pb.DefaultCreateUser(ctx, req.Payload, s.db)
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserResponse{Result: user}, err

}
func (s *Server) Read(ctx context.Context, req *pb.ReadUserRequest) (*pb.ReadUserResponse, error) {
	return &pb.ReadUserResponse{Result: &pb.User{Id: req.Id}}, nil
}

func (s *Server) Update(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	return &pb.UpdateUserResponse{}, nil
}

func (s *Server) List(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserResponse, error) {
	return &pb.ListUserResponse{}, nil
}

func (s *Server) Delete(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	return &pb.DeleteUserResponse{}, nil
}

func init() {
	config = cfg.NewViperConfig()
	if config.GetBool(`debug`) {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.Debug("Service RUN on DEBUG mode")
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}
	resource.RegisterApplication(config.GetString(`app.id`))
	resource.SetPlural()

	logrus.Debugf("resource application name: %v", resource.ApplicationName())
}

func main() {
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
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		logrus.Fatal(err)
	}
	defer db.Close()
	db.LogMode(true)
	db.SetLogger(logrus.NewEntry(logrus.New()))
	db.DB().SetMaxOpenConns(1)

	server := NewServer(db)

	server.db.AutoMigrate(&pb.UserORM{})

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, server)
	defer grpcServer.Stop()

	listen, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
	defer listen.Close()

	err = grpcServer.Serve(listen)
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}
