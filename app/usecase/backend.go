package usecase

import (
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/reviewsys/backend/app/domain/model"
	"github.com/reviewsys/backend/app/domain/repository"
)

type BackendUsecase interface {
	GetVersions(*empty.Empty) (*model.Versions, error)
}

type backendUsecase struct {
	repo repository.BackendRepository
}

func NewBackendUsecase(repo repository.BackendRepository) *backendUsecase {
	return &backendUsecase{
		repo: repo,
	}
}

func (u *backendUsecase) GetVersions(e *empty.Empty) (*model.Versions, error) {
	return u.repo.GetVersions(e)
}
