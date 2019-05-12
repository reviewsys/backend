package usecase

import (
	"github.com/infobloxopen/atlas-app-toolkit/rpc/resource"
	"github.com/reviewsys/backend/app/domain/model"
	"github.com/reviewsys/backend/app/domain/repository"
	"github.com/reviewsys/backend/app/domain/service"
)

type UserUsecase interface {
	GetByID(id *resource.Identifier) (*model.User, error)
	Update(*model.User) error
	GetByTeamID(teamID int64) (*model.User, error)
	Store(*model.User) error
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

func (u *userUsecase) Store(user *model.User) error {
	return u.repo.Store(user)
}

func (u *userUsecase) Delete(user *model.User) error {
	return u.repo.Delete(user)
}
