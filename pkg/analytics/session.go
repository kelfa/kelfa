package analytics

import (
	"time"

	"go.kelfa.io/kelfa/pkg/session"
)

func (a *Analytics) GetSessions(mit time.Duration) *session.Sessions {
	var ss session.Sessions
	for _, s := range a.dataPoints {
		s := s
		ss.AddDataPoint(&s)
	}
	ss.SplitSessions(mit)
	return &ss
}
