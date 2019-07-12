package analytics

import "go.kelfa.io/kelfa/pkg/dal/objects"

func (a *Analytics) CacheDataAvailable() bool {
	if len(a.dataPoints[0].CDNCache) == 0 {
		return false
	}
	return true
}

func (a *Analytics) CacheStats() *objects.CacheStats {
	cs := objects.CacheStats{}
	for _, log := range a.dataPoints {
		cs.Total = cs.Total + 1
		switch log.CDNCache {
		case "Hit":
			cs.Hits = cs.Hits + 1
		case "RefreshHit":
			cs.RefreshHits = cs.RefreshHits + 1
		case "Miss":
			cs.Misses = cs.Misses + 1
		case "LimitExceeded":
			cs.LimitExceeded = cs.LimitExceeded + 1
		case "CapacityExceeded":
			cs.CapacityExceeded = cs.CapacityExceeded + 1
		case "Error":
			cs.Errors = cs.Errors + 1
		case "Redirect":
			cs.Redirects = cs.Redirects + 1
		default:
			cs.Others = cs.Others + 1
		}
	}
	return &cs
}
