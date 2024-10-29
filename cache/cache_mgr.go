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
	if cm.cache != nil {
		cm.cache.Close()
	}
}

func (cm *CacheMgr) Get(key string) ([]byte, error) {
	if cm.cache == nil {
		return nil, fmt.Errorf("cache not found")
	}

	value, ok := cm.cache.Get(key)
	if !ok {
		return nil, fmt.Errorf("failed to get key: %s from cache", key)
	}
	return value, nil
}

func (cm *CacheMgr) Set(key string, value []byte) error {
	if cm.cache == nil {
		return fmt.Errorf("cache not found")
	}

	if ok := cm.cache.Set(key, value); !ok {
		return fmt.Errorf("failed to set cache, key: %s", key)
	}
	return nil
}

func (cm *CacheMgr) GetCacheStats() (*CacheStats, error) {
	if cm.cache == nil {
		return &CacheStats{}, fmt.Errorf("cache not found")
	}

	stats := cm.cache.Stats()
	entries := cm.cache.Size()
	return &CacheStats{
		hits:    stats.Hits(),
		misses:  stats.Misses(),
		ratio:   stats.Ratio(),
		entries: entries,
	}, nil
}
