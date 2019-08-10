package analytics

import (
	"fmt"
	"time"
)

type Mode int

const (
	None Mode = iota
	Hourly
	Daily
	Weekly
	Monthly
)

func ParseMode(s string) (Mode, error) {
	switch s {
	case "none":
		return None, nil
	case "hourly":
		return Hourly, nil
	case "daily":
		return Daily, nil
	case "weekly":
		return Weekly, nil
	case "monthly":
		return Monthly, nil
	default:
		return None, fmt.Errorf("the '%v' mode is not valid. Valid modes are: none, hourly, daily, weekly, monthly", s)
	}
}

type TimeSlot struct {
	Begin time.Time
	End   time.Time
}

func createSlots(from time.Time, to time.Time, m Mode) (ts []TimeSlot) {
	switch m {
	case 0:
		ts = append(ts, TimeSlot{
			Begin: from,
			End:   to,
		})
	case 1:
		pointer := time.Date(from.Year(), from.Month(), from.Day(), from.Hour(), 0, 0, 0, from.Location())
		for {
			if pointer.After(to) {
				break
			}
			ts = append(ts, TimeSlot{
				Begin: pointer,
				End:   pointer.Add(time.Hour).Add(-time.Second),
			})
			pointer = pointer.Add(time.Hour)
		}
	case 2:
		pointer := time.Date(from.Year(), from.Month(), from.Day(), 0, 0, 0, 0, from.Location())
		for {
			if pointer.After(to) {
				break
			}
			ts = append(ts, TimeSlot{
				Begin: pointer,
				End:   pointer.AddDate(0, 0, 1).Add(-time.Second),
			})
			pointer = pointer.AddDate(0, 0, 1)
		}
	case 3:
		pointer := time.Date(from.Year(), from.Month(), from.Day(), 0, 0, 0, 0, from.Location())
		for {
			if pointer.After(to) {
				break
			}
			ts = append(ts, TimeSlot{
				Begin: pointer,
				End:   pointer.AddDate(0, 0, 7).Add(-time.Second),
			})
			pointer = pointer.AddDate(0, 0, 7)
		}
	case 4:
		pointer := time.Date(from.Year(), from.Month(), 0, 0, 0, 0, 0, from.Location())
		for {
			if pointer.After(to) {
				break
			}
			ts = append(ts, TimeSlot{
				Begin: pointer,
				End:   pointer.AddDate(0, 1, 0).Add(-time.Second),
			})
			pointer = pointer.AddDate(0, 1, 0)
		}
	}
	return ts
}
