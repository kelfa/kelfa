package session

import (
	"fmt"
	"sort"
	"time"
)

// Split a session to multiple sessions based on a given maximum inactive time
func SplitSessionsByMaxInactiveTime(s Session, mit time.Duration) (ss []Session) {
	var pdpdt time.Time
	lastCut := 0
	sort.Slice(s.DataPoints, func(i, j int) bool {
		return s.DataPoints[i].DateTime.Before(s.DataPoints[j].DateTime)
	})
	for k, dp := range s.DataPoints {
		if k == 0 {
			pdpdt = dp.DateTime
		}
		if pdpdt.Add(mit).Before(dp.DateTime) {
			// Ignoring error, safe by definition
			ns, err := NewSessionWithDataPoints(s.DataPoints[lastCut:k])
			if err != nil {
				fmt.Println(err)
			}
			ss = append(ss, *ns)
			lastCut = k
		}
	}
	if lastCut != len(s.DataPoints)-1 {
		ns, _ := NewSessionWithDataPoints(s.DataPoints[lastCut : len(s.DataPoints)-1])
		ss = append(ss, *ns)
	}
	return ss
}
