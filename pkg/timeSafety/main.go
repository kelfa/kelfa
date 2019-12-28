package timeSafety

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

func IsTimeSlotDaySafe(timeslot string) (bool, error) {
	t, err := time.ParseInLocation("2006-01-02-15", timeslot, time.UTC)
	if err != nil {
		return false, err
	}
	td := time.Date(t.Year(), t.Month(), t.Day()+1, 0, 0, 0, 0, time.UTC)
	safeTime := time.Now().In(time.UTC).Add(-25 * time.Hour)
	if safeTime.Before(td) {
		return false, nil
	}
	return true, nil
}
