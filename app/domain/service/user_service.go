package service

import (
	"github.com/pkg/errors"

	"github.com/infobloxopen/atlas-app-toolkit/rpc/resource"
	"github.com/reviewsys/backend/app/domain/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Duplicated(id *resource.Identifier) error {
	user, err := s.repo.GetByID(id)
	if user != nil {
		return errors.Errorf("%v already exists", id)
	}
	if err != nil {
		return err
	}
	return nil
}
