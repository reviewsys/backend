package registry

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/reviewsys/backend/app/domain/model"
	"github.com/reviewsys/backend/app/domain/service"
	"github.com/reviewsys/backend/app/interface/persistence/database"
	"github.com/reviewsys/backend/app/usecase"
	"github.com/sarulabs/di"
	log "github.com/sirupsen/logrus"
)

type Container struct {
	ctn di.Container
}

func NewContainer(dsn string) (*Container, error) {
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
		{
			Name: "postgres",
			Build: func(ctn di.Container) (interface{}, error) {
				db, err := gorm.Open("postgres", dsn)
				db.DB().SetMaxOpenConns(1)
				return db, err
			},
			Close: func(obj interface{}) error {
				return obj.(*gorm.DB).Close()
			},
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
	// Retrieve the connection.
	db := ctn.Get("postgres").(*gorm.DB)
	db.AutoMigrate(&model.User{})
	repo := database.NewUserRepository(db)
	service := service.NewUserService(repo)
	return usecase.NewUserUsecase(repo, service), nil
}
