package reliableget
import "net/http"

type Config struct {
	Cache Cache `json:cache`
}

type Cache struct {
	Engine string `json:"engine"`
	Url string `json:"url"`
}

type Options struct {
	Timeout int
	CacheKey string
	CacheTTL int
	ExplicitNoCache bool
	Headers string
	Tracer string
	Type string
	StatsdKey string
	EventHandler func(error, *http.Response)
}

type ReliableGet struct {
	config Config
}

func NewReliableGet(config *Config) *ReliableGet {
	m := &ReliableGet{}
	m.config = config
	return m
}

func (r *ReliableGet) Get(Url string, options *Options) {

}

func (r *ReliableGet) Disconnect() {

}

func pipeAndCacheContent() {

}

func getWithNoCache() {

}