package rpc

import (
	"github.com/reviewsys/backend/app/interface/rpc/v1.0"
	"github.com/reviewsys/backend/app/registry"
	"google.golang.org/grpc"
)

func Apply(server *grpc.Server, ctn *registry.Container) {
	v1.Apply(server, ctn)
}
