package analytics

import (
	"errors"
	"fmt"
	"time"

	"go.kelfa.io/kelfa/pkg/dal"
	"go.kelfa.io/kelfa/pkg/session"
)

type Mode int

const (
	None Mode = iota
	Hourly
	Daily
	Weekly
	Monthly
)

type DataSet struct {
	Begin    time.Time
	End      time.Time
	Sessions session.Sessions
}

type Analytics struct {
	Begin        time.Time
	End          time.Time
	GroupingMode Mode
	Data         []DataSet
}

// TODO: Split Data based on modes. For now period always = None
func New(ds *dal.DataSource, from time.Time, to time.Time, mode Mode, mit time.Duration) (*Analytics, error) {
	dsbt, err := (*ds).DataBeginTime()
	if err != nil {
		return nil, err
	}
	if from.Before(*dsbt) {
		return nil, fmt.Errorf("the data start at %v so %v is an invalid starting point", dsbt, from)
	}
	dset, err := (*ds).DataEndTime()
	if err != nil {
		return nil, err
	}
	if to.After(*dset) {
		return nil, fmt.Errorf("the data end at %v so %v is an invalid ending point", dset, to)
	}
	dsdps, err := (*ds).GetDataPoints(from, to)
	if err != nil {
		return nil, err
	}

	var ss session.Sessions
	for _, s := range dsdps {
		s := s
		ss.AddDataPoint(&s)
	}
	ss.SplitSessions(mit)

	if len(ss.Sessions) == 0 {
		return nil, errors.New("no data to show")
	}

	a := Analytics{
		Begin:        from,
		End:          to,
		GroupingMode: mode,
		Data: []DataSet{
			{
				Begin:    from,
				End:      to,
				Sessions: ss,
			},
		},
	}

	return &a, nil
}
