package memory

import (
	"github.com/golang/protobuf/ptypes/empty"
)

const version = "0.0.1"

type backendRepository struct {
	Version string
}

func NewBackendRepository() *backendRepository {
	return &backendRepository{Version: version}
}

func (r *backendRepository) GetVersion(e *empty.Empty) (string, error) {
	return r.Version, nil
}
