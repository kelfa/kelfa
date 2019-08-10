package analytics

import (
	"testing"
	"time"
)

func TestCreateSlots(t *testing.T) {
	tests := []struct {
		From       time.Time
		To         time.Time
		Mode       Mode
		ExpectedTS []TimeSlot
	}{
		{ // 3 standard slots
			From: time.Date(2019, time.July, 01, 0, 0, 0, 0, time.UTC),
			To:   time.Date(2019, time.July, 01, 2, 59, 59, 0, time.UTC),
			Mode: Hourly,
			ExpectedTS: []TimeSlot{
				{
					Begin: time.Date(2019, time.July, 01, 0, 0, 0, 0, time.UTC),
					End:   time.Date(2019, time.July, 01, 0, 59, 59, 0, time.UTC),
				},
				{
					Begin: time.Date(2019, time.July, 01, 1, 0, 0, 0, time.UTC),
					End:   time.Date(2019, time.July, 01, 1, 59, 59, 0, time.UTC),
				},
				{
					Begin: time.Date(2019, time.July, 01, 2, 0, 0, 0, time.UTC),
					End:   time.Date(2019, time.July, 01, 2, 59, 59, 0, time.UTC),
				},
			},
		},
		{ //  Alter min
			From: time.Date(2019, time.July, 01, 0, 23, 0, 0, time.UTC),
			To:   time.Date(2019, time.July, 01, 0, 59, 59, 0, time.UTC),
			Mode: Hourly,
			ExpectedTS: []TimeSlot{
				{
					Begin: time.Date(2019, time.July, 01, 0, 0, 0, 0, time.UTC),
					End:   time.Date(2019, time.July, 01, 0, 59, 59, 0, time.UTC),
				},
			},
		},
		{ //  Alter max
			From: time.Date(2019, time.July, 01, 0, 0, 0, 0, time.UTC),
			To:   time.Date(2019, time.July, 01, 0, 23, 0, 0, time.UTC),
			Mode: Hourly,
			ExpectedTS: []TimeSlot{
				{
					Begin: time.Date(2019, time.July, 01, 0, 0, 0, 0, time.UTC),
					End:   time.Date(2019, time.July, 01, 0, 59, 59, 0, time.UTC),
				},
			},
		},
	}
	for _, tc := range tests {
		tss := createSlots(tc.From, tc.To, tc.Mode)
		if len(tss) != len(tc.ExpectedTS) {
			t.Errorf("counted %v timeslots instead of %v", len(tss), len(tc.ExpectedTS))
		}
	}
}
