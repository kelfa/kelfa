package main

import "time"

func IsTimeSlotSafe(timeslot string) (bool, error) {
	t, err := time.ParseInLocation("2006-01-02-15", timeslot, time.UTC)
	if err != nil {
		return false, err
	}
	safeTime := time.Now().In(time.UTC).Add(-25 * time.Hour)
	if safeTime.Before(t) {
		return false, nil
	}
	return true, nil
}
