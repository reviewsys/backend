package repository

import (
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/reviewsys/backend/app/domain/model"
)

type BackendRepository interface {
	GetVersions(*empty.Empty) (*model.Versions, error)
}
