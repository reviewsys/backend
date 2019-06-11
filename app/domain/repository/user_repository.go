package repository

import (
	"context"

	"github.com/infobloxopen/atlas-app-toolkit/rpc/resource"
	"github.com/reviewsys/backend/app/domain/model"
	pb "github.com/reviewsys/backend/app/interface/rpc/v1.0/protocol"
)

type UserRepository interface {
	GetByID(id *resource.Identifier) (*model.User, error)
	Update(*model.User) error
	GetByTeamID(teamID int64) (*model.User, error)
	Store(context.Context, *pb.User) (*pb.User, error)
	Delete(*model.User) error
}
