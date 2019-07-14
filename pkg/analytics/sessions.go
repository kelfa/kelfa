package analytics

import (
	"fmt"
	"net"
	"time"

	"go.kelfa.io/kelfa/pkg/dal/objects"
)

// The idea is that a session is a single user (IP + UserAgent) in a single day, which might not be correct.
type Session struct {
	Date       time.Time
	IP         net.IP
	UserAgent  string
	DataPoints []objects.DataPoint
}

func NewSession(dp *objects.DataPoint) Session {
	return Session{
		Date:       time.Date(dp.DateTime.Year(), dp.DateTime.Month(), dp.DateTime.Day(), 0, 0, 0, 0, time.UTC),
		IP:         dp.ClientIP,
		UserAgent:  dp.ClientUserAgent,
		DataPoints: []objects.DataPoint{*dp},
	}
}

func (s *Session) AddDataPoint(dp *objects.DataPoint) error {
	if !dp.ClientIP.Equal(s.IP) {
		return fmt.Errorf("impossible to add a datapoint that is referring to %v to a session about %v", dp.ClientIP, s.IP)
	}
	if dp.ClientUserAgent != s.UserAgent {
		return fmt.Errorf("impossible to add a datapoint that is referring to %s to a session about %s", dp.ClientUserAgent, s.UserAgent)
	}
	if !s.Date.Equal(time.Date(dp.DateTime.Year(), dp.DateTime.Month(), dp.DateTime.Day(), 0, 0, 0, 0, time.UTC)) {
		return fmt.Errorf("impossible to add a datapoint that is referring to %v to a session about %v", time.Date(dp.DateTime.Year(), dp.DateTime.Month(), dp.DateTime.Day(), 0, 0, 0, 0, time.UTC), s.Date)
	}
	s.DataPoints = append(s.DataPoints, *dp)
	return nil
}

type Sessions struct {
	Sessions []Session
}

func (ss *Sessions) AddDataPoint(dp *objects.DataPoint) {
	for k, s := range ss.Sessions {
		if dp.ClientIP.Equal(s.IP) &&
			dp.ClientUserAgent == s.UserAgent &&
			s.Date.Equal(time.Date(dp.DateTime.Year(), dp.DateTime.Month(), dp.DateTime.Day(), 0, 0, 0, 0, time.UTC)) {
			// Ignoring the error since we have just checked that no error will occur
			_ = (*ss).Sessions[k].AddDataPoint(dp)
			return
		}
	}
	ns := NewSession(dp)
	ss.Sessions = append(ss.Sessions, ns)
}

func (a *Analytics) GetSessions() *Sessions {
	var ss Sessions
	for _, s := range a.dataPoints {
		ss.AddDataPoint(&s)
	}
	return &ss
}
