package dal

import (
	"fmt"
	"time"

	"go.kelfa.io/kelfa/pkg/dal/filesystem"
	"go.kelfa.io/kelfa/pkg/dal/objects"
	"go.kelfa.io/kelfa/pkg/dal/sql"
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
	case "sqlite":
		return sql.New(bo)
	}
	return nil, fmt.Errorf("unable to find any back-end called %s", backend)
}

type WritableDataSource interface {
	DataSource
	AddDataPoint(objects.DataPoint) error
	AddDataPoints([]objects.DataPoint) error
}

func NewWritableDataSource(backend string, bo objects.BackendOptions) (WritableDataSource, error) {
	switch backend {
	case "sqlite":
		return sql.New(bo)
	}
	return nil, fmt.Errorf("unable to find any back-end called %s", backend)
}
