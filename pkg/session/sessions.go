package session

import (
	"go.kelfa.io/kelfa/pkg/dal/objects"
)

type Sessions struct {
	Sessions []Session
}

func (ss *Sessions) SplitSessions() {
	// session_inactivity_timeout
}

func (ss *Sessions) AddDataPoint(dp *objects.DataPoint) {
	for k, s := range ss.Sessions {
		if dp.ClientIP.Equal(*s.IP) &&
			dp.ClientUserAgent == *s.UserAgent {
			// Ignoring the error since we have just checked that no error will occur
			_ = (*ss).Sessions[k].AddDataPoint(dp)
			return
		}
	}
	ns := Session{}
	ns.AddDataPoint(dp)
	ss.Sessions = append(ss.Sessions, ns)
}
