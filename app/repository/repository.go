package repository

import (
	"github.com/infobloxopen/atlas-app-toolkit/rpc/resource"
	"github.com/reviewsys/backend/app/domain/model"
)

type UserRepository interface {
	GetByID(id *resource.Identifier) (*model.User, error)
	Update(*model.User) error
	GetByTeamID(teamID int64) (*model.User, error)
	Store(*model.User) error
	Delete(*model.User) error
}
