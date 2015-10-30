package reliableget
import (
	"github.com/caledhwa/gongeal/reliableget/cache"
	"github.com/caledhwa/gongeal/reliableget/config"
)

type ReliableGet struct {
	config *config.Config
	factory *cache.CacheFactory
	cache cache.CacheEngine
}

func New(config *config.Config) *ReliableGet {
	g := &ReliableGet{}
	g.config = config
	g.factory = cache.NewCacheFactory()
	g.cache = g.factory.GetCache(&config.Cache)
	return g
}

func (r *ReliableGet) Get(Url string, options *config.Options) {

}

func (r *ReliableGet) Disconnect() {

}

func pipeAndCacheContent() {

}

func getWithNoCache() {

}