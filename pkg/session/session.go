package session

import (
	"errors"
	"fmt"
	"net"
	"sort"
	"time"

	"go.kelfa.io/kelfa/pkg/dal/objects"
)

type Session struct {
	Begin      *time.Time
	End        *time.Time
	IP         *net.IP
	UserAgent  *string
	DataPoints []objects.DataPoint
}

func (s *Session) isValidDataPoint(dp *objects.DataPoint) error {
	if dp == nil {
		return errors.New("impossible to add a nil data-point to a session")
	}
	if s.IP != nil && !s.IP.Equal(dp.ClientIP) {
		return fmt.Errorf("impossible to add a data-point that is referring to %v to a session about %v", dp.ClientIP, s.IP)
	}
	if s.UserAgent != nil && *s.UserAgent != dp.ClientUserAgent {
		return fmt.Errorf("impossible to add a data-point that is referring to %s to a session about %v", dp.ClientUserAgent, s.UserAgent)
	}
	return nil
}

func (s *Session) syncTimes() {
	sort.Slice(s.DataPoints[:], func(i, j int) bool {
		return s.DataPoints[i].DateTime.Before(s.DataPoints[j].DateTime)
	})
	s.Begin = &s.DataPoints[0].DateTime
	s.End = &s.DataPoints[len(s.DataPoints)-1].DateTime
}

func (s *Session) AddDataPoint(dp *objects.DataPoint) error {
	if err := s.isValidDataPoint(dp); err != nil {
		return err
	}
	s.DataPoints = append(s.DataPoints, *dp)
	if s.IP == nil {
		s.IP = &dp.ClientIP
	}
	if s.UserAgent == nil {
		s.UserAgent = &dp.ClientUserAgent
	}
	s.syncTimes()
	return nil
}

// TODO: Make this stronger against the injection of data-points from different IP/UA
func NewSessionWithDataPoints(dps []objects.DataPoint) (*Session, error) {
	s := Session{}
	for _, dp := range dps {
		s.DataPoints = append(s.DataPoints, dp)
	}
	if s.IP == nil {
		s.IP = &dps[0].ClientIP
	}
	if s.UserAgent == nil {
		s.UserAgent = &dps[0].ClientUserAgent
	}
	s.syncTimes()
	return &s, nil
}
