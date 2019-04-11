package repository

import "github.com/reviewsys/backend/app/models"

type UserRepository interface {
	Fetch(cursor string, num int64) ([]*models.User, error)
	GetByID(id int64) (*models.User, error)
	Update(user *models.User) (*models.User, error)
	GetByTeamID(teamID int64) (*models.User, error)
	Store(a *models.User) (int64, error)
	Delete(id int64) (bool, error)
}
