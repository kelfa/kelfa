package dal

import (
	"fmt"
	"time"

	"go.kelfa.io/kelfa/pkg/dal/filesystem"
	"go.kelfa.io/kelfa/pkg/dal/objects"
)

type DataSource interface {
	DataBeginTime() (*time.Time, error)
	DataEndTime() (*time.Time, error)
	GetDataPoints(from time.Time, to time.Time) ([]objects.DataPoint, error)
}

func New(backend string, bo objects.BackendOptions) (DataSource, error) {
	switch backend {
	case "filesystem":
		return filesystem.New(bo)
	}
	return nil, fmt.Errorf("unable to find any back-end called %s", backend)
}
