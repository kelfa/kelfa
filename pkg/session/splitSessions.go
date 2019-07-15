package session

import "time"

// Split a session to multiple sessions based on a given maximum inactive time
func SplitSessionsByMaxInactiveTime(s Session, mit time.Duration) (ss []Session) {
	var pdpdt time.Time
	lastCut := 0
	for k, dp := range s.DataPoints {
		if k == 0 {
			pdpdt = dp.DateTime
		}
		if pdpdt.Add(mit).Before(dp.DateTime) {
			// Ignoring error, safe by definition
			ns, _ := NewSessionWithDataPoints(s.DataPoints[lastCut:k])
			ss = append(ss, *ns)
		}
	}
	ns, _ := NewSessionWithDataPoints(s.DataPoints[lastCut : len(s.DataPoints)-1])
	ss = append(ss, *ns)
	return ss
}
