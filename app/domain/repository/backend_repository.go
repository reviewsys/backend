package repository

import "github.com/golang/protobuf/ptypes/empty"

type BackendRepository interface {
	GetVersion(e *empty.Empty) (string, error)
}
