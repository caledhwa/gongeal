package cache
import "github.com/caledhwa/gongeal/reliableget/config"

type MemoryCache struct {
	config *config.Cache
	engine string
}

func NewMemoryCache(_config *config.Cache) CacheEngine {
	c := &MemoryCache{}
	c.config = _config
	c.engine = "memorycache"
	return c
}
func (c *MemoryCache) Get(key string, next func(e error, cacheData string, oldCacheData string)) {

}

func (c *MemoryCache) Set(key string, value string, ttl int, next func()) {

}

func (c *MemoryCache) Engine() string {
	return c.engine
}