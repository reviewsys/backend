package v1

import (
	"context"

	"github.com/reviewsys/backend/app/interface/rpc/v1.0/protocol"
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

func (s *userService) Create(ctx context.Context, req *protocol.CreateUserRequest) (*protocol.CreateUserResponse, error) {
	return &protocol.CreateUserResponse{}, nil
}

func (s *userService) Read(ctx context.Context, req *protocol.ReadUserRequest) (*protocol.ReadUserResponse, error) {
	return &protocol.ReadUserResponse{}, nil
}

func (s *userService) Update(ctx context.Context, req *protocol.UpdateUserRequest) (*protocol.UpdateUserResponse, error) {
	return &protocol.UpdateUserResponse{}, nil
}

func (s *userService) List(ctx context.Context, req *protocol.ListUserRequest) (*protocol.ListUserResponse, error) {
	return &protocol.ListUserResponse{}, nil
}

func (s *userService) Delete(ctx context.Context, req *protocol.DeleteUserRequest) (*protocol.DeleteUserResponse, error) {
	return &protocol.DeleteUserResponse{}, nil
}
