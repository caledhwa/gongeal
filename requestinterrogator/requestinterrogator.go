package requestinterrogator

import (
	"github.com/caledhwa/gongeal/config"
	"github.com/caledhwa/gongeal/util"
	"net/http"
	"net/url"
)

type RequestInterrogator struct {
	Configuration *config.Config
}

func NewRequestInterrogator(configuration *config.Config) (*RequestInterrogator)  {

	return &RequestInterrogator{Configuration:configuration}
}

func (interrogator *RequestInterrogator) InterrogateRequest(request *http.Request) (map[string]string) {
	params :=  make(map[string]string)

	queryParams := interrogator.interrogateParams(request.URL.Query())
	pageUrl := getPageUrl(request)

	for key, value := range queryParams {
		params["param:" + key] = value
	}

	params["url:href"] = pageUrl
	encodedUrl, _ := util.EncodeUrl(pageUrl)
	params["url:href:encoded"] = encodedUrl

	return params
}

func getPageUrl(request *http.Request) string {

	var pageUrl string

	// Check host headers for hostname
	host := request.Header.Get("http_host")
	if host == "" {
		host = request.Header.Get("host")
	}

	// Grab protocol and default to http if none available
	protocol := request.URL.Scheme
	if protocol == "" {
		protocol = "http"
	}

	search := request.URL.RawQuery
	if search != "" {
		search = "?" + search
	}
	pathname := request.URL.Path[0:]

	pageUrl = protocol + "://" + host + pathname + search

	return pageUrl
}

func (interrogator *RequestInterrogator) interrogateParams(params url.Values) map[string]string {
	returnParams := make(map[string]string)
	if interrogator.Configuration.Query != nil {
		for _, item := range interrogator.Configuration.Query {
			queryItem := params.Get(item.Key)
			if queryItem != "" {
				returnParams[item.MapTo] = queryItem
			}
		}
	}
	return returnParams
}
