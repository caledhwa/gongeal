package config
import "net/http"

type Query struct {
	Key string `json:"key"`
	MapTo string `json:"mapTo"`
}

type Url struct {
	Names []string `json:"names"`
	Pattern string   `json:"pattern"`
}

type Parameters struct {
	Query []Query `json:"query"`

	Servers struct {
			  Local string `json:"local"`
		  } `json:"servers"`

	Urls []Url `json:"urls"`
}

type Cdn struct {
	URL string `json:"url"`
}

type Backend struct {
	Pattern string `json:"pattern"`
	Target string `json:"target"`
	Host string `json:"host"`
	Timeout string `json:"timeout"`
	CacheKey string `json:"cacheKey"`
	TTL string `json:"ttl"`
	QuietFailure *bool `json:"quietFailure"`
	LeaveContentOnFail *bool `json:"leaveContentOnFail"`
	DontPassUrl *bool `json:"dontPassUrl"`
	PassThrough bool `json:"passThrough"`
	ReplaceOuter *bool  `json:"replaceOuter"`
	ContentTypes []string `json:"contentTypes"`
	Fn string `json:"fn"`
}

type BackendSelectorFunction func (*http.Request, map[string]string) bool
type StatusCodeHandlerFunction func(*http.Request, http.ResponseWriter, map[string]string, interface{} )

type Config struct {

	SelectorFunctions map[string]BackendSelectorFunction `json:"selectorFunctions"`

	Backend []Backend `json:"backend"`

	Cache struct {
	   	Engine string `json:"engine"`
	} `json:"cache"`

	Cdn Cdn `json:"cdn"`

	Circuitbreaker struct {
	   ErrorThreshold  int `json:"errorThreshold"`
	   NumBuckets      int `json:"numBuckets"`
	   VolumeThreshold int `json:"volumeThreshold"`
	   WindowDuration  int `json:"windowDuration"`
	} `json:"circuitbreaker"`

	Cookies struct {
		Whitelist []string `json:"whitelist"`
	} `json:"cookies"`

	FollowRedirect bool `json:"followRedirect"`

	Parameters Parameters `json:"parameters"`

	StatusCodeHandlers struct {

		_02 struct {
			Fn string `json:"fn"`
		} `json:"302"`

		_03 struct {
			Data struct {
				Redirect string `json:"redirect"`
			} `json:"data"`
			Fn string `json:"fn"`
		} `json:"403"`

	} `json:"statusCodeHandlers"`
}
