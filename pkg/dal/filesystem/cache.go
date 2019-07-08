package filesystem

import (
	"go.kelfa.io/kelfa/dal/objects"
)

func (d *DP) CacheDataAvailable() bool {
	if _, ok := d.data[0]["x-edge-response-result-type"]; ok {
		return true
	}
	return false
}

func (d *DP) CacheStats() *objects.CacheStats {
	cs := objects.CacheStats{}
	for _, log := range d.data {
		cs.Total = cs.Total + 1
		switch log["x-edge-response-result-type"] {
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
