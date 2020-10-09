package big_cache

import (
	"github.com/allegro/bigcache"
	"time"
)

func MakeBigCache() *bigcache.BigCache {
	cache, _ := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))
	return cache
}
