package registry

import (
	"context"
	"database/sql"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/reviewsys/backend/app/domain/service"
	"github.com/reviewsys/backend/app/interface/persistence/database"
	"github.com/reviewsys/backend/app/interface/persistence/memory"
	pb "github.com/reviewsys/backend/app/interface/rpc/v1.0/protocol"
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
			Scope: di.App,
			Build: buildUserUsecase,
		},
		{
			Name:  "backend-usecase",
			Scope: di.App,
			Build: buildBackendUsecase,
		},
		{
			Name:  "postgres-pool",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				db, err := gorm.Open("postgres", dsn)
				db.LogMode(true)
				db.SetLogger(log.NewEntry(log.New()))
				db.DB().SetMaxOpenConns(1)
				return db, err
			},
			Close: func(obj interface{}) error {
				return obj.(*gorm.DB).Close()
			},
		},
		{
			// Define the connection in the Request scope.
			// Each request will use its own connection.
			Name:  "postgres",
			Scope: di.Request,
			Build: func(ctn di.Container) (interface{}, error) {
				pool := ctn.Get("postgres-pool").(*sql.DB)
				return pool.Conn(context.Background())
			},
			Close: func(obj interface{}) error {
				return obj.(*sql.Conn).Close()
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

func (c *Container) Delete() error {
	return c.ctn.Delete()
}

func buildUserUsecase(ctn di.Container) (interface{}, error) {
	// Retrieve the connection.
	db := ctn.Get("postgres-pool").(*gorm.DB)
	db.AutoMigrate(&pb.UserORM{})
	repo := database.NewUserRepository(db)
	service := service.NewUserService(repo)
	return usecase.NewUserUsecase(repo, service), nil
}

func buildBackendUsecase(ctn di.Container) (interface{}, error) {
	repo := memory.NewBackendRepository()
	return usecase.NewBackendUsecase(repo), nil
}
