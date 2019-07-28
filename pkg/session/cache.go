package session

type CacheStats struct {
	Hits             int
	RefreshHits      int
	Misses           int
	LimitExceeded    int
	CapacityExceeded int
	Errors           int
	Redirects        int
	Total            int
	Others           int
}

func (c *CacheStats) Add(i *CacheStats) {
	c.Hits = c.Hits + i.Hits
	c.RefreshHits = c.RefreshHits + i.RefreshHits
	c.Misses = c.Misses + i.Misses
	c.LimitExceeded = c.LimitExceeded + i.LimitExceeded
	c.CapacityExceeded = c.CapacityExceeded + i.CapacityExceeded
	c.Errors = c.Errors + i.Errors
	c.Redirects = c.Redirects + i.Redirects
	c.Total = c.Total + i.Total
	c.Others = c.Others + i.Others
}

func (s *Session) CacheStats() *CacheStats {
	cs := CacheStats{}
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
	return &cs
}
