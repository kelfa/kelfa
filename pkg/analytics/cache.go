package analytics

import "go.kelfa.io/kelfa/pkg/dal/objects"

func (a *Analytics) CacheDataAvailable() bool {
	return len(a.dataPoints[0].CDNCache) == 0
}

func (a *Analytics) CacheStats() *objects.CacheStats {
	cs := objects.CacheStats{}
	for _, log := range a.dataPoints {
		cs.Total++
		switch log.CDNCache {
		case "Hit":
			cs.Hits++
		case "RefreshHit":
			cs.RefreshHits++
		case "Miss":
			cs.Misses++
		case "LimitExceeded":
			cs.LimitExceeded++
		case "CapacityExceeded":
			cs.CapacityExceeded++
		case "Error":
			cs.Errors++
		case "Redirect":
			cs.Redirects++
		default:
			cs.Others++
		}
	}
	return &cs
}
