package cache

type CacheStats struct {
	hits    int64
	misses  int64
	ratio   float64
	entries int
}

func (cs *CacheStats) GetHits() int64 {
	return cs.hits
}

func (cs *CacheStats) GetMisses() int64 {
	return cs.misses
}

func (cs *CacheStats) GetRatio() float64 {
	return cs.ratio
}

func (cs *CacheStats) GetEntries() int {
	return cs.entries
}
