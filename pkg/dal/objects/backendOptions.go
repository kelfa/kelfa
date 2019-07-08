package objects

import "time"

type BackendOptions struct {
	Path string
	From time.Time
	To   time.Time
}
