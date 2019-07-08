package objects

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
