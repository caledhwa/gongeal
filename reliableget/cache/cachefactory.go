package cache

import "github.com/caledhwa/gongeal/reliableget/config"

type CacheEngine interface {
	Get(key string, next func(e error, cacheData string, oldCacheData string))
	Set(key string, value string, ttl int, next func())
	Engine() string
}

type CacheFactory struct {
	cacheInstance map[string]CacheEngine
}

var Engines = map[string](func(*config.Cache) CacheEngine) {
	"redis" : NewRedisCache,
	"nocache" : NewNoCache,
	"memorycache" : NewMemoryCache,
}

func NewCacheFactory() *CacheFactory {
	c := &CacheFactory{}
	c.cacheInstance = make(map[string]CacheEngine)
	return c
}

func (c *CacheFactory) GetCache(_config *config.Cache) CacheEngine {

	if _config == nil || _config.Engine == "" || Engines[_config.Engine] == nil {
		_config = &config.Cache{ Engine: "nocache" }
	}

	if c.cacheInstance[_config.Engine] == nil {
		c.cacheInstance[_config.Engine] = Engines[_config.Engine](_config)
	}

	return c.cacheInstance[_config.Engine]
}

func (c *CacheFactory) ClearCacheInstance() {
	c.cacheInstance = make(map[string]CacheEngine)
}