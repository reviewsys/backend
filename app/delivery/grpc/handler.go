package grpc

import (
	log "github.com/sirupsen/logrus"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"context"

	"github.com/infobloxopen/atlas-app-toolkit/rpc/resource"
	"github.com/reviewsys/backend/app/delivery/grpc/user"
	"github.com/reviewsys/backend/app/domain/model"
	_usecase "github.com/reviewsys/backend/app/usecase"
)

func NewAppServerGrpc(gserver *grpc.Server, userUcase _usecase.UserUsecase) {

	appServer := &server{
		usecase: userUcase,
	}

	user.RegisterUserServiceServer(gserver, appServer)
	reflection.Register(gserver)
}

type server struct {
	usecase _usecase.UserUsecase
}

func (s *server) Read(ctx context.Context, req *user.ReadUserRequest) (*user.ReadUserResponse, error) {
	var id *resource.Identifier
	if req != nil {
		id = req.Id
	}
	resp, err := s.usecase.GetByID(id)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if resp == nil {
		return nil, model.NOT_FOUND_ERROR
	}
	return &user.ReadUserResponse(resp), nil
}
