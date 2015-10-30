package cache
import "github.com/caledhwa/gongeal/reliableget/config"

type NoCache struct {
	config *config.Cache
	engine string
}

func NewNoCache(_config *config.Cache) CacheEngine {
	c := &NoCache{}
	c.config = _config
	c.engine = "nocache"
	return c
}

func (c *NoCache) Get(key string, next func(e error, cacheData string, oldCacheData string)) {
}

func (c *NoCache) Set(key string, value string, ttl int, next func()) {
}

func (c *NoCache) Engine() string {
	return c.engine
}