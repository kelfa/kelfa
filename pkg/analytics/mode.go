package analytics

import (
	"fmt"
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
