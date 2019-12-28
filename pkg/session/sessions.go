package session

import (
	"time"

	"go.kelfa.io/pkg/dal/objects"
)

type Sessions []Session

func (ss *Sessions) SplitSessions(mit time.Duration) {
	nss := []Session{}
	for _, s := range *ss {
		nss = append(nss, SplitSessionsByMaxInactiveTime(s, mit)...)
	}
	*ss = nss
}

func (ss *Sessions) AddDataPoint(dp *objects.DataPoint) {
	for k, s := range *ss {
		if dp.ClientIP.Equal(*s.IP) &&
			dp.ClientUserAgent == *s.UserAgent {
			// Ignoring the error since we have just checked that no error will occur
			_ = (*ss)[k].AddDataPoint(dp)
			return
		}
	}
	ns := Session{}
	_ = ns.AddDataPoint(dp) // Surely safe, being the first one
	*ss = append(*ss, ns)
}
