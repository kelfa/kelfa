package filesystem

import (
	"time"

	"github.com/kelfa/kelfa/pkg/session"
)

func (d *DP) DailySessions(from time.Time, to time.Time) (map[time.Time][]session.Session, error) {
	return nil, nil
}
