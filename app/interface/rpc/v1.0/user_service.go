package v1

import (
	"context"

	pb "github.com/reviewsys/backend/app/interface/rpc/v1.0/protocol"
	"github.com/reviewsys/backend/app/usecase"
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
	user, err := s.userUsecase.Store(ctx, req.Payload)
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserResponse{Result: user}, err

}

func (s *userService) Read(ctx context.Context, req *pb.ReadUserRequest) (*pb.ReadUserResponse, error) {
	return &pb.ReadUserResponse{Result: &pb.User{Id: req.Id}}, nil
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
