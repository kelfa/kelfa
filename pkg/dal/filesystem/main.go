package filesystem

import (
	"go.kelfa.io/pkg/dal/objects"
)

type DP struct {
	folder string
}

// TODO: Check if folder exists
func New(bo objects.BackendOptions) (*DP, error) {
	d := DP{
		folder: bo.Path,
	}
	return &d, nil
}
