package usecase

import (
	"strconv"
	"time"

	"github.com/reviewsys/backend/app/models"
	"github.com/reviewsys/backend/app/repository"
)

type UserUsecase interface {
	Fetch(cursor string, num int64) ([]*models.User, string, error)
	GetByID(id int64) (*models.User, error)
	Update(ar *models.User) (*models.User, error)
	GetByTeamID(teamID int64) (*models.User, error)
	Store(*models.User) (*models.User, error)
	Delete(id int64) (bool, error)
}

type userUsecase struct {
	userRepos repository.UserRepository
}

func (u *userUsecase) Fetch(cursor string, num int64) ([]*models.User, string, error) {
	if num == 0 {
		num = 10
	}

	listUser, err := u.userRepos.Fetch(cursor, num)
	if err != nil {
		return nil, "", err
	}
	nextCursor := ""

	if size := len(listUser); size == int(num) {
		lastID := listUser[num-1].ID
		nextCursor = strconv.Itoa(int(lastID))
	}

	return listUser, nextCursor, nil
}

func (u *userUsecase) GetByID(id int64) (*models.User, error) {

	return u.userRepos.GetByID(id)
}

func (u *userUsecase) Update(m *models.User) (*models.User, error) {
	_, err := u.userRepos.GetByID(m.ID)
	if err != nil {
		return nil, err
	}

	m.UpdatedAt = time.Now()
	return u.userRepos.Update(m)
}

func (u *userUsecase) GetByTeamID(teamID int64) (*models.User, error) {

	return u.userRepos.GetByTeamID(teamID)
}

func (u *userUsecase) Store(m *models.User) (*models.User, error) {

	existedUser, _ := u.GetByTeamID(m.TeamID)
	if existedUser != nil {
		return nil, models.CONFLIT_ERROR
	}

	id, err := u.userRepos.Store(m)
	if err != nil {
		return nil, err
	}

	m.ID = id
	return m, nil
}

func (u *userUsecase) Delete(id int64) (bool, error) {
	existedUser, _ := u.GetByID(id)

	if existedUser == nil {
		return false, models.NOT_FOUND_ERROR
	}

	return u.userRepos.Delete(id)
}

func NewUserUsecase(u repository.UserRepository) UserUsecase {
	return &userUsecase{u}
}
