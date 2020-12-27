package sql

import (
	"time"

	"go.kelfa.io/pkg/session"
)

func (d *DP) DailySessions(from time.Time, to time.Time) (map[time.Time][]session.Session, error) {
	var items []session.Session
	if err := d.conn.Where("begin >= ?", from).Where("end <= ?", to).Find(&items).Error; err != nil {
		return nil, err
	}
	sessions := make(map[time.Time][]session.Session)
	for _, item := range items {
		day := time.Date(item.Begin.Year(), item.Begin.Month(), item.Begin.Day(), 0, 0, 0, 0, time.UTC)
		if sessions[day] == nil {
			sessions[day] = []session.Session{item}
		} else {
			sessions[day] = append(sessions[day], item)
		}
	}
	return sessions, nil
}
