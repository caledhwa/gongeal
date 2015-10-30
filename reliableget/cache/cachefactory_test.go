package cache_test
import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
	"github.com/caledhwa/gongeal/reliableget/config"
	"github.com/caledhwa/gongeal/reliableget/cache"
)

func TestCacheFactory(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TestCacheFactory")
}

var _ = Describe("Cache Factory", func() {

	It("should return redis cache engine when specified in config", func() {
		cacheFactory := cache.NewCacheFactory()
		engine := cacheFactory.GetCache(&config.Cache{
			Engine:"redis",
		}).(cache.CacheEngine)
		Expect(engine.Engine()).To(BeEquivalentTo("redis"))
	})

	It("should return memory cache engine when specified in config", func() {
		cacheFactory := cache.NewCacheFactory()
		engine := cacheFactory.GetCache(&config.Cache{
			Engine:"memorycache",
		}).(cache.CacheEngine)
		Expect(engine.Engine()).To(BeEquivalentTo("memorycache"))
	})

	It("should return nocache engine when no config specified", func() {
		cacheFactory := cache.NewCacheFactory()
		engine := cacheFactory.GetCache(nil).(cache.CacheEngine)
		Expect(engine.Engine()).To(BeEquivalentTo("nocache"))
	})

	It("should return nocache engine when config specified with no Engine set", func() {
		cacheFactory := cache.NewCacheFactory()
		engine := cacheFactory.GetCache(&config.Cache{}).(cache.CacheEngine)
		Expect(engine.Engine()).To(BeEquivalentTo("nocache"))
	})

	It("should return nocache engine when config specified with emptystring Engine set", func() {
		cacheFactory := cache.NewCacheFactory()
		engine := cacheFactory.GetCache(&config.Cache{ Engine: "" }).(cache.CacheEngine)
		Expect(engine.Engine()).To(BeEquivalentTo("nocache"))
	})

	It("should return nocache engine when engine not found", func() {
		cacheFactory := cache.NewCacheFactory()
		engine := cacheFactory.GetCache(&config.Cache{ Engine: "joebob" }).(cache.CacheEngine)
		Expect(engine.Engine()).To(BeEquivalentTo("nocache"))
	})

})
