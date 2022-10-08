package session_test

import (
	"net"
	"testing"
	"time"

	"github.com/kelfa/kelfa/pkg/dal/objects"
	"github.com/kelfa/kelfa/pkg/session"
)

func TestSession(t *testing.T) {
	tests := []struct {
		Datapoints        []objects.DataPoint
		ExpectedBegin     time.Time
		ExpectedEnd       time.Time
		ExpectedIP        net.IP
		ExpectedUserAgent string
	}{
		{
			Datapoints: []objects.DataPoint{
				{
					DateTime:        time.Date(2019, time.July, 01, 0, 0, 0, 0, time.UTC),
					ClientIP:        net.IPv4(8, 8, 8, 8),
					ClientUserAgent: "test",
				},
			},
			ExpectedBegin:     time.Date(2019, time.July, 01, 0, 0, 0, 0, time.UTC),
			ExpectedEnd:       time.Date(2019, time.July, 01, 0, 0, 0, 0, time.UTC),
			ExpectedIP:        net.IPv4(8, 8, 8, 8),
			ExpectedUserAgent: "test",
		},
		{
			Datapoints: []objects.DataPoint{
				{
					DateTime:        time.Date(2019, time.July, 01, 3, 0, 0, 0, time.UTC),
					ClientIP:        net.IPv4(8, 8, 8, 8),
					ClientUserAgent: "test",
				},
				{
					DateTime:        time.Date(2019, time.July, 01, 0, 0, 0, 0, time.UTC),
					ClientIP:        net.IPv4(8, 8, 8, 8),
					ClientUserAgent: "test",
				},
			},
			ExpectedBegin:     time.Date(2019, time.July, 01, 0, 0, 0, 0, time.UTC),
			ExpectedEnd:       time.Date(2019, time.July, 01, 3, 0, 0, 0, time.UTC),
			ExpectedIP:        net.IPv4(8, 8, 8, 8),
			ExpectedUserAgent: "test",
		},
	}

	for _, st := range tests {
		s := session.Session{}
		for c, dp := range st.Datapoints {
			dp := dp
			err := s.AddDataPoint(&dp)
			if err != nil {
				t.Errorf("an error occurred while adding the datapoint #%v object: %v", c, err)
			}
		}
		if !s.Begin.Equal(st.ExpectedBegin) {
			t.Errorf("the expected begin was %v while the actual is %v", st.ExpectedBegin, s.Begin)
		}
		if !s.End.Equal(st.ExpectedEnd) {
			t.Errorf("the expected end was %v while the actual is %v", st.ExpectedEnd, s.End)
		}
		if !s.IP.Equal(st.ExpectedIP) {
			t.Errorf("the expected IP was %v while the actual is %v", st.ExpectedIP, s.IP)
		}
		if *s.UserAgent != st.ExpectedUserAgent {
			t.Errorf("the expected User Agent was %s while the actual is %v", st.ExpectedUserAgent, s.UserAgent)
		}
	}
}
