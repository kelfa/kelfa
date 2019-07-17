package analytics

import "go.kelfa.io/kelfa/pkg/session"

func (a *Analytics) GetSessions() *session.Sessions {
	var ss session.Sessions
	for _, s := range a.dataPoints {
		s := s
		ss.AddDataPoint(&s)
	}
	return &ss
}
