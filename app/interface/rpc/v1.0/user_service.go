package v1

import (
	"context"

	pb "github.com/reviewsys/backend/app/interface/rpc/v1.0/protocol"
	"github.com/reviewsys/backend/app/usecase"
	log "github.com/sirupsen/logrus"
)

type userService struct {
	userUsecase usecase.UserUsecase
}

func NewUserService(userUsecase usecase.UserUsecase) *userService {
	return &userService{
		userUsecase: userUsecase,
	}
}

func (s *userService) Create(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	return &pb.CreateUserResponse{}, nil
}

func (s *userService) Read(ctx context.Context, req *pb.ReadUserRequest) (*pb.ReadUserResponse, error) {
	log.Info("Read Request: ", req)
	return &pb.ReadUserResponse{}, nil
}

func (s *userService) Update(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	return &pb.UpdateUserResponse{}, nil
}

func (s *userService) List(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserResponse, error) {
	return &pb.ListUserResponse{}, nil
}

func (s *userService) Delete(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	return &pb.DeleteUserResponse{}, nil
}
