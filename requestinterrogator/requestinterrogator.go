package requestinterrogator

import (
	"github.com/caledhwa/gongeal/config"
	"net/http"
)

type RequestInterrogator struct {
	Configuration *config.Config
}

func NewRequestInterrogator(configuration *config.Config) (*RequestInterrogator)  {
	return &RequestInterrogator{}
}

func (interrogator *RequestInterrogator) InterrogateRequest(request *http.Request) (map[string]string) {
	return make(map[string]string)
}