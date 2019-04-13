package repository

import "github.com/reviewsys/backend/app/models"

type UserRepository interface {
	GetByID(id int64) (*models.User, error)
	Update(*models.User) error
	GetByTeamID(teamID int64) (*models.User, error)
	Store(*models.User) error
	Delete(*models.User) error
}
