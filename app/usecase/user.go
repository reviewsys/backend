package usecase

import (
	"github.com/reviewsys/backend/app/domain/model"
	"github.com/reviewsys/backend/app/repository"
	"github.com/infobloxopen/atlas-app-toolkit/rpc/resource"
)

type UserUsecase interface {
	GetByID(id *resource.Identifier) (*model.User, error)
	Update(*model.User) error
	GetByTeamID(teamID int64) (*model.User, error)
	Store(*model.User) error
	Delete(*model.User) error
}

type userUsecase struct {
	userRepos repository.UserRepository
}

func (u *userUsecase) GetByID(id *resource.Identifier) (*model.User, error) {
	return u.userRepos.GetByID(id)
}

func (u *userUsecase) Update(user *model.User) error {
	return u.userRepos.Update(user)
}

func (u *userUsecase) GetByTeamID(teamID int64) (*model.User, error) {
	return u.userRepos.GetByTeamID(teamID)
}

func (u *userUsecase) Store(user *model.User) error {
	return u.userRepos.Store(user)
}

func (u *userUsecase) Delete(user *model.User) error {
	return u.userRepos.Delete(user)
}

func NewUserUsecase(u repository.UserRepository) UserUsecase {
	return &userUsecase{u}
}
