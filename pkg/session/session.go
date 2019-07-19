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
		return fmt.Errorf("mismatching ip: expecting '%v', received '%v'", s.IP, dp.ClientIP)
	}
	if s.UserAgent != nil && *s.UserAgent != dp.ClientUserAgent {
		return fmt.Errorf("mismatching user-agent: expecting '%v', received '%s'", s.UserAgent, dp.ClientUserAgent)
	}
	return nil
}

func (s *Session) syncTimes() {
	sort.Slice(s.DataPoints, func(i, j int) bool {
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

func (s *Session) Validate() error {
	if s.DataPoints == nil {
		return errors.New("session without any data-point")
	}
	if s.IP == nil {
		return errors.New("ip is null")
	}
	if s.UserAgent == nil {
		return errors.New("user agent is null")
	}
	for k, dp := range s.DataPoints {
		if dp.ClientIP.Equal(*s.IP) {
			return fmt.Errorf("expecting the '%v' IP, while the '%s' was found in line %v", s.IP, dp.ClientIP, k)
		}
		if *s.UserAgent != dp.ClientUserAgent {
			return fmt.Errorf("expecting the '%v' UA, while the '%s' was found in line %v", s.UserAgent, dp.ClientUserAgent, k)
		}
	}
	return nil
}

func NewSessionWithDataPoints(dps []objects.DataPoint) (*Session, error) {
	s := Session{}
	s.DataPoints = append(s.DataPoints, dps...)
	s.IP = &dps[0].ClientIP
	s.UserAgent = &dps[0].ClientUserAgent
	s.syncTimes()
	if err := s.Validate(); err != nil {
		return nil, err
	}
	return &s, nil
}
