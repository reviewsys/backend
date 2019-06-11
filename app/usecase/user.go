package usecase

import (
	"context"

	"github.com/infobloxopen/atlas-app-toolkit/rpc/resource"
	"github.com/reviewsys/backend/app/domain/model"
	"github.com/reviewsys/backend/app/domain/repository"
	"github.com/reviewsys/backend/app/domain/service"
	pb "github.com/reviewsys/backend/app/interface/rpc/v1.0/protocol"
)

type UserUsecase interface {
	GetByID(id *resource.Identifier) (*model.User, error)
	Update(*model.User) error
	GetByTeamID(teamID int64) (*model.User, error)
	Store(context.Context, *pb.User) (*pb.User, error)
	Delete(*model.User) error
}

type userUsecase struct {
	repo    repository.UserRepository
	service *service.UserService
}

func NewUserUsecase(repo repository.UserRepository, service *service.UserService) *userUsecase {
	return &userUsecase{
		repo:    repo,
		service: service,
	}
}

func (u *userUsecase) GetByID(id *resource.Identifier) (*model.User, error) {
	return u.repo.GetByID(id)
}

func (u *userUsecase) Update(user *model.User) error {
	return u.repo.Update(user)
}

func (u *userUsecase) GetByTeamID(teamID int64) (*model.User, error) {
	return u.repo.GetByTeamID(teamID)
}

func (u *userUsecase) Store(ctx context.Context, user *pb.User) (*pb.User, error) {
	return u.repo.Store(ctx, user)
}

func (u *userUsecase) Delete(user *model.User) error {
	return u.repo.Delete(user)
}
