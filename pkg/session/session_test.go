package session_test

import (
	"net"
	"testing"
	"time"

	"go.kelfa.io/kelfa/pkg/dal/objects"
	"go.kelfa.io/kelfa/pkg/session"
)

type SessionAddDataPointsTest struct {
	Datapoints        []objects.DataPoint
	ExpectedBegin     time.Time
	ExpectedEnd       time.Time
	ExpectedIP        net.IP
	ExpectedUserAgent string
}

var sessionAddDataPointsTests = []SessionAddDataPointsTest{
	SessionAddDataPointsTest{
		Datapoints: []objects.DataPoint{
			objects.DataPoint{
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
	SessionAddDataPointsTest{
		Datapoints: []objects.DataPoint{
			objects.DataPoint{
				DateTime:        time.Date(2019, time.July, 01, 3, 0, 0, 0, time.UTC),
				ClientIP:        net.IPv4(8, 8, 8, 8),
				ClientUserAgent: "test",
			},
			objects.DataPoint{
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

func TestCacheStats(t *testing.T) {
	for _, st := range sessionAddDataPointsTests {
		s := session.Session{}
		for c, dp := range st.Datapoints {
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
