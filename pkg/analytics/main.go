package analytics

import (
	"fmt"
	"time"

	"go.kelfa.io/kelfa/pkg/dal"
	"go.kelfa.io/kelfa/pkg/dal/objects"
)

type Analytics struct {
	from       time.Time
	to         time.Time
	dataPoints []objects.DataPoint
}

func New(ds *dal.DataSource, from time.Time, to time.Time) (*Analytics, error) {
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
	a := Analytics{
		from:       from,
		to:         to,
		dataPoints: dsdps,
	}
	return &a, nil
}
