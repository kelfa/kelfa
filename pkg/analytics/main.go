package analytics

import (
	"errors"
	"fmt"
	"time"

	"go.kelfa.io/kelfa/pkg/dal"
	"go.kelfa.io/kelfa/pkg/session"
)

type DataSets []DataSet

type DataSet struct {
	Begin    time.Time
	End      time.Time
	Sessions session.Sessions
}

func GenerateDataSets(from time.Time, to time.Time, m Mode) *DataSets {
	var dss DataSets
	slots := createSlots(from, to, m)
	for _, slot := range slots {
		dss = append(dss, DataSet{
			Begin: slot.Begin,
			End:   slot.End,
		})
	}
	return &dss
}

func (dss *DataSets) AddSession(s *session.Session) {
}

type Analytics struct {
	Begin        time.Time
	End          time.Time
	GroupingMode Mode
	Sessions     session.Sessions
}

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

	if len(ss) == 0 {
		return nil, errors.New("no data to show")
	}

	a := Analytics{
		Begin:        from,
		End:          to,
		GroupingMode: mode,
		Sessions:     ss,
	}

	return &a, nil
}
