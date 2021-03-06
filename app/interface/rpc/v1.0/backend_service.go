package v1

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/reviewsys/backend/app/interface/rpc/v1.0/protocol"
	"github.com/reviewsys/backend/app/usecase"
)

type backendService struct {
	backendUsecase usecase.BackendUsecase
}

func NewBackendService(backendUsecase usecase.BackendUsecase) *backendService {
	return &backendService{
		backendUsecase: backendUsecase,
	}
}

func (s *backendService) GetVersions(ctx context.Context, e *empty.Empty) (*pb.VersionResponse, error) {
	v, err := s.backendUsecase.GetVersions(e)
	return &pb.VersionResponse{
		Version:   v.Version,
		Revision:  v.Revision,
		BuildDate: v.BuildDate,
	}, err
}
