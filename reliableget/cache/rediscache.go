package cache
import "github.com/caledhwa/gongeal/reliableget/config"

type RedisCache struct {
	config *config.Cache
	engine string
}

func NewRedisCache(_config *config.Cache) CacheEngine {
	c := &RedisCache{}
	c.config = _config
	c.engine = "redis"
	return c
}
func (c *RedisCache) Get(key string, next func(e error, cacheData string, oldCacheData string)) {

}

func (c *RedisCache) Set(key string, value string, ttl int, next func()) {

}

func (c *RedisCache) Engine() string {
	return c.engine
}