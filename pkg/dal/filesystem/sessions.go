package filesystem

import (
	"time"

	"go.kelfa.io/pkg/session"
)

func (d *DP) DailySessions(from time.Time, to time.Time) (map[time.Time][]session.Session, error) {
	return nil, nil
}
