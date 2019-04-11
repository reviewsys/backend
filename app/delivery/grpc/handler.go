package grpc

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"context"

	google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
	"github.com/reviewsys/backend/app/delivery/grpc/user"
	models "github.com/reviewsys/backend/app/models"
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

func (s *server) transformUserRPC(u *models.User) *user.User {

	if u == nil {
		return nil
	}

	updatedAt := &google_protobuf.Timestamp{

		Seconds: u.UpdatedAt.Unix(),
	}
	createdAt := &google_protobuf.Timestamp{
		Seconds: u.CreatedAt.Unix(),
	}
	res := &user.User{
		ID:        u.ID,
		TeamId:    u.TeamID,
		Name:      u.Name,
		IsAdmin:   u.IsAdmin,
		UpdatedAt: updatedAt,
		CreatedAt: createdAt,
	}
	return res
}

func (s *server) GetUser(ctx context.Context, in *user.SingleRequest) (*user.User, error) {
	id := int64(0)
	if in != nil {
		id = in.Id
	}
	u, err := s.usecase.GetByID(id)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	if u == nil {
		return nil, models.NOT_FOUND_ERROR
	}

	res := s.transformUserRPC(u)
	return res, nil
}
