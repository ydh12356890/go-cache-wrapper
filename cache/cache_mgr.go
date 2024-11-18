package cache

import (
	"fmt"
	"time"

	"github.com/maypok86/otter"
)

type CacheMgr struct {
	cache *otter.Cache[string, []byte]
}

func NewCacheMgr(capacity int, ttl int) (*CacheMgr, error) {
	cache, err := otter.MustBuilder[string, []byte](capacity).
		CollectStats().
		Cost(func(key string, value []byte) uint32 {
			return uint32(len(key) + len(value))
		}).
		WithTTL(time.Duration(ttl) * time.Second).Build()
	if err != nil {
		return nil, fmt.Errorf("failed to new cache mgr, capacity: %d, ttl: %d, err: %v", capacity, ttl, err)
	}

	return &CacheMgr{
		cache: &cache,
	}, nil
}

func (cm *CacheMgr) Stop() {
	cm.cache.Close()
}

func (cm *CacheMgr) Has(key string) bool {
	return cm.cache.Has(key)
}

func (cm *CacheMgr) Delete(key string) {
	cm.cache.Delete(key)
}

func (cm *CacheMgr) Get(key string) ([]byte, bool) {
	return cm.cache.Get(key)
}

func (cm *CacheMgr) Set(key string, value []byte) bool {
	return cm.cache.Set(key, value)
}

func (cm *CacheMgr) GetCacheStats() *CacheStats {
	stats := cm.cache.Stats()
	entries := cm.cache.Size()
	return &CacheStats{
		hits:    stats.Hits(),
		misses:  stats.Misses(),
		ratio:   stats.Ratio(),
		entries: entries,
	}
}
