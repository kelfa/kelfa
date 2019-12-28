package filesystem_test

import (
	"testing"
	"time"

	"go.kelfa.io/pkg/dal/filesystem"
	"go.kelfa.io/pkg/dal/objects"
)

func TestDataEndTime(t *testing.T) {
	tests := []struct {
		folder       string
		expectedTime time.Time
	}{
		{
			folder:       "../../../test/data/logs",
			expectedTime: time.Date(2019, time.July, 1, 23, 59, 59, 0, time.UTC),
		},
	}
	for _, test := range tests {
		f, err := filesystem.New(objects.BackendOptions{Path: test.folder})
		if err != nil {
			t.Errorf("an error occurred while creating the filesystem object: %v", err)
		}
		bt, err := f.DataEndTime()
		if err != nil {
			t.Errorf("an error occurred while looking for the begin time: %v", err)
		}
		if *bt != test.expectedTime {
			t.Errorf("expecting %v, received %v", test.expectedTime, bt)
		}
	}
}
