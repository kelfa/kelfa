package analytics_test

import (
	"testing"
	"time"

	"go.kelfa.io/kelfa/pkg/analytics"
	"go.kelfa.io/kelfa/pkg/dal"
	"go.kelfa.io/kelfa/pkg/dal/objects"
)

type CacheTest struct {
	From   time.Time
	To     time.Time
	Hits   int
	Misses int
}

var cacheTestData = []CacheTest{
	CacheTest{
		From:   time.Date(2019, time.July, 01, 0, 0, 0, 0, time.UTC),
		To:     time.Date(2019, time.July, 01, 23, 59, 59, 0, time.UTC),
		Hits:   423,
		Misses: 253,
	},
}

func TestCacheStats(t *testing.T) {
	data, err := dal.New("filesystem", objects.BackendOptions{Path: "../../data"})
	if err != nil {
		t.Errorf("an error occurred while creating the filesystem object: %v", err)
	}
	for _, dp := range cacheTestData {
		a, err := analytics.New(&data, dp.From, dp.To)
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
