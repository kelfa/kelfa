package timeSafety_test

import (
	"testing"
	"time"

	"go.kelfa.io/aws-cloudfront-logCompactor/pkg/timeSafety"
)

type dataPoint struct {
	Time time.Time
	Safe bool
}

func TestIsTimeSlotSafe(t *testing.T) {

	dps := []dataPoint{
		{
			Time: time.Now().In(time.UTC).Add(-26 * time.Hour),
			Safe: true,
		},
		{
			Time: time.Now().In(time.UTC).Add(-24 * time.Hour),
			Safe: false,
		},
	}

	for _, dp := range dps {
		act, err := timeSafety.IsTimeSlotSafe(dp.Time.Format("2006-01-02-15"))
		if err != nil {
			t.Fatal(err)
		}
		if act != dp.Safe {
			t.Errorf("%v was expected to be %v to safety while it was %v", dp.Time, dp.Safe, act)
		}
	}
}

func TestIsTimeSlotDaySafe(t *testing.T) {

	dps := []dataPoint{
		{
			Time: time.Now().In(time.UTC).Add(-48 * time.Hour),
			Safe: true,
		},
		{
			Time: time.Now().In(time.UTC).Add(-24 * time.Hour),
			Safe: false,
		},
	}

	for _, dp := range dps {
		act, err := timeSafety.IsTimeSlotDaySafe(dp.Time.Format("2006-01-02-15"))
		if err != nil {
			t.Fatal(err)
		}
		if act != dp.Safe {
			t.Errorf("%v was expected to be %v to safety while it was %v", dp.Time, dp.Safe, act)
		}
	}
}
