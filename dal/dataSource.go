package dal

import (
	"time"

	"go.kelfa.io/kelfa/dal/objects"
)

type DataSource interface {
	GetBeginDate() (*time.Time, error)
	GetEndDate() (*time.Time, error)

	CacheDataAvailable() bool
	CacheStats() *objects.CacheStats
}
