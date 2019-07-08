package dal

import (
	"fmt"

	"go.kelfa.io/kelfa/dal/filesystem"
	"go.kelfa.io/kelfa/dal/objects"
)

func New(backend string, bo objects.BackendOptions) (DataSource, error) {
	switch backend {
	case "filesystem":
		return filesystem.New(bo)
	}
	return nil, fmt.Errorf("unable to find any backend called %s", backend)
}
