package v1

import (
	"github.com/reviewsys/backend/app/interface/rpc/v1.0/protocol"
	"github.com/reviewsys/backend/app/registry"
	"github.com/reviewsys/backend/app/usecase"
	"google.golang.org/grpc"
)

func Apply(server *grpc.Server, ctn *registry.Container) {
	protocol.RegisterUserServiceServer(server, NewUserService(ctn.Resolve("user-usecase").(usecase.UserUsecase)))
	protocol.RegisterBackendServer(server, NewBackendService(ctn.Resolve("backend-usecase").(usecase.BackendUsecase)))
}
