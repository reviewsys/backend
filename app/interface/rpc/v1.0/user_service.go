package v1

import (
	log "github.com/sirupsen/logrus"

	"context"

	"github.com/infobloxopen/atlas-app-toolkit/rpc/resource"
	"github.com/reviewsys/backend/app/domain/model"
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
	var id *resource.Identifier
	if req != nil {
		id = req.Id
	}
	resp, err := s.userUsecase.GetByID(id)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if resp == nil {
		return nil, model.NOT_FOUND_ERROR
	}
	return &protocol.ReadUserResponse{
		Result: &protocol.User{
			Id:      resp.ID,
			TeamId:  resp.TeamID,
			IsAdmin: resp.IsAdmin,
		},
	}, nil
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
