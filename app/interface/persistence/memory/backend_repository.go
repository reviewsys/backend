package memory

import (
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/reviewsys/backend/app/domain/model"
)

const (
	Version   = "0.0.2"
	Revision  = "0000000"
	BuildDate = "2006-01-02T04:05:06Z"
)

type backendRepository struct {
	version   string
	revision  string
	buildDate string
}

func NewBackendRepository() *backendRepository {
	return &backendRepository{
		version:   Version,
		revision:  Revision,
		buildDate: BuildDate,
	}
}

func (r *backendRepository) GetVersions(e *empty.Empty) (*model.Versions, error) {
	return &model.Versions{
		Version:   r.version,
		Revision:  r.revision,
		BuildDate: r.buildDate,
	}, nil
}
