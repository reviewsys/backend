package registry

import (
	"github.com/reviewsys/backend/app/domain/service"
	"github.com/reviewsys/backend/app/interface/persistence/database"
	"github.com/reviewsys/backend/app/usecase"
	"github.com/sarulabs/di"
	log "github.com/sirupsen/logrus"
)

type Container struct {
	ctn di.Container
}

func NewContainer() (*Container, error) {
	builder, err := di.NewBuilder()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if err := builder.Add([]di.Def{
		{
			Name:  "user-usecase",
			Build: buildUserUsecase,
		},
	}...); err != nil {
		log.Error(err)
		return nil, err
	}

	return &Container{
		ctn: builder.Build(),
	}, nil
}

func (c *Container) Resolve(name string) interface{} {
	return c.ctn.Get(name)
}

func (c *Container) Clean() error {
	return c.ctn.Clean()
}

func buildUserUsecase(ctn di.Container) (interface{}, error) {
	repo := database.NewUserRepository()
	service := service.NewUserService(repo)
	return usecase.NewUserUsecase(repo, service), nil
}
