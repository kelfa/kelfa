package analytics_test

import (
	"testing"
	"time"

	"go.kelfa.io/pkg/analytics"
	"go.kelfa.io/pkg/dal"
	"go.kelfa.io/pkg/dal/objects"
)

func TestCacheStats(t *testing.T) {
	tests := []struct {
		From   time.Time
		To     time.Time
		Hits   int
		Misses int
	}{
		{
			From:   time.Date(2019, time.July, 01, 0, 0, 0, 0, time.UTC),
			To:     time.Date(2019, time.July, 01, 23, 59, 59, 0, time.UTC),
			Hits:   9,
			Misses: 10,
		},
	}

	data, err := dal.NewDataSource("filesystem", objects.BackendOptions{Path: "../../test/data/logs"})
	if err != nil {
		t.Errorf("an error occurred while creating the filesystem object: %v", err)
	}
	for _, dp := range tests {
		a, err := analytics.New(&data, dp.From, dp.To, analytics.None, time.Hour)
		if err != nil {
			t.Errorf("an error occurred while creating the analytics object: %v", err)
		}
		stats := a.CacheStats()
		if stats.Hits != dp.Hits {
			t.Errorf("counted %v hits instead of %v", stats.Hits, dp.Hits)
		}
		if stats.Misses != dp.Misses {
			t.Errorf("counted %v misses instead of %v", stats.Misses, dp.Misses)
		}
	}
}
