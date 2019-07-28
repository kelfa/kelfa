package analytics

import "go.kelfa.io/kelfa/pkg/dal/objects"

func (a *Analytics) CacheDataAvailable() bool {
	return len(a.Data[0].Sessions.Sessions[0].DataPoints[0].CDNCache) == 0
}

func (a *Analytics) CacheStats() *objects.CacheStats {
	cs := objects.CacheStats{}
	for _, ds := range a.Data {
		for _, s := range ds.Sessions.Sessions {
			for _, log := range s.DataPoints {
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
		}
	}
	return &cs
}
