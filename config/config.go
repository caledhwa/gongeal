package config

type Query struct {
	Key string `json:"key"`
	MapTo string `json:"mapTo"`
}

type Config struct {

	Backend []struct {
		Pattern string `json:"pattern"`
		Target  string `json:"target"`
	} `json:"backend"`

	Cache struct {
	   	Engine string `json:"engine"`
	} `json:"cache"`

	Cdn struct {
	URL string `json:"url"`
	} `json:"cdn"`

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

	Parameters struct {

		Servers struct {
			Local string `json:"local"`
		} `json:"servers"`

		Urls []struct {
			Names []string `json:"names"`
			Pattern string   `json:"pattern"`
		} `json:"urls"`

	} `json:"parameters"`

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

	Query []Query `json:"query"`
}
