package session_test

import (
	"testing"
	"time"

	"go.kelfa.io/pkg/dal"
	"go.kelfa.io/pkg/dal/objects"
	"go.kelfa.io/pkg/session"
)

func TestSessions(t *testing.T) {
	tests := []struct {
		From     time.Time
		To       time.Time
		Sessions int
	}{
		{
			From:     time.Date(2019, time.July, 01, 0, 0, 0, 0, time.UTC),
			To:       time.Date(2019, time.July, 01, 23, 59, 59, 0, time.UTC),
			Sessions: 5,
		},
	}

	data, err := dal.NewDataSource("filesystem", objects.BackendOptions{Path: "../../test/data/logs"})
	if err != nil {
		t.Errorf("an error occurred while creating the file-system object: %v", err)
	}
	for _, dp := range tests {
		dsdps, err := data.GetDataPoints(dp.From, dp.To)
		if err != nil {
			t.Errorf("an error occurred while retrieving the data points: %v", err)
		}

		var ss session.Sessions
		for _, s := range dsdps {
			s := s
			ss.AddDataPoint(&s)
		}
		ss.SplitSessions(time.Hour)

		if len(ss) != dp.Sessions {
			t.Errorf("counted %v sessions instead of %v", len(ss), dp.Sessions)
		}
	}
}
