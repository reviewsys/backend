package usecase

import (
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/reviewsys/backend/app/domain/repository"
)

type BackendUsecase interface {
	GetVersion(*empty.Empty) (string, error)
}

type backendUsecase struct {
	repo repository.BackendRepository
}

func NewBackendUsecase(repo repository.BackendRepository) *backendUsecase {
	return &backendUsecase{
		repo: repo,
	}
}

func (u *backendUsecase) GetVersion(e *empty.Empty) (string, error) {
	return u.repo.GetVersion(e)
}
