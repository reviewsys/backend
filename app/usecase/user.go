package usecase

import (
	"github.com/reviewsys/backend/app/models"
	"github.com/reviewsys/backend/app/repository"
)

type UserUsecase interface {
	GetByID(id int64) (*models.User, error)
	Update(*models.User) error
	GetByTeamID(teamID int64) (*models.User, error)
	Store(*models.User) error
	Delete(*models.User) error
}

type userUsecase struct {
	userRepos repository.UserRepository
}

func (u *userUsecase) GetByID(id int64) (*models.User, error) {
	return u.userRepos.GetByID(id)
}

func (u *userUsecase) Update(user *models.User) error {
	return u.userRepos.Update(user)
}

func (u *userUsecase) GetByTeamID(teamID int64) (*models.User, error) {
	return u.userRepos.GetByTeamID(teamID)
}

func (u *userUsecase) Store(user *models.User) error {
	return u.userRepos.Store(user)
}

func (u *userUsecase) Delete(user *models.User) error {
	return u.userRepos.Delete(user)
}

func NewUserUsecase(u repository.UserRepository) UserUsecase {
	return &userUsecase{u}
}
