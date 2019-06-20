package main

import (
	"testing"
	"time"
)

//func IsTimeSlotSafe(timeslot string) (bool, error) {
func TestIsTimeSlotSafe(T *testing.T) {

	type DataPoint struct {
		Time time.Time
		Safe bool
	}

	dps := []DataPoint{
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
		act, err := IsTimeSlotSafe(dp.Time.Format("2006-01-02-15"))
		if err != nil {
			T.Fatal(err)
		}
		if act != dp.Safe {
			T.Errorf("%v was expected to be %v to safety while it was %v", dp.Time, dp.Safe, act)
		}
	}
}
