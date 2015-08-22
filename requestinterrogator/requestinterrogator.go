package requestinterrogator

import (
	"github.com/caledhwa/gongeal/config"
	"github.com/caledhwa/gongeal/util"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

type RequestInterrogator struct {
	configParams *config.Parameters
	configCdn *config.Cdn
}

func NewRequestInterrogator(configParams *config.Parameters, configCdn *config.Cdn) (*RequestInterrogator)  {

	return &RequestInterrogator{configParams:configParams, configCdn:configCdn}
}

func (interrogator *RequestInterrogator) InterrogateRequest(request *http.Request) (map[string]string) {
	params :=  make(map[string]string)

	queryParams := interrogator.interrogateParams(request.URL.Query())
	templateParams := interrogator.interrogatePath(request.URL.Path)
	pageUrl := getPageUrl(request)

	for key, value := range queryParams {
		params["param:" + key] = value
	}

	for key, value := range templateParams {
		params["param:" + key] = value
	}

	for key, value := range request.URL.Query() {
		queryValues := value[0]
		for _,val := range value[1:] {
			queryValues = queryValues + "," + val
		}
		params["query:" + key] = queryValues
	}

	for key, value := range request.Header {
		headerValues := value[0]
		for _,val := range value[1:] {
			headerValues = headerValues + "," + val
		}
		params["header:" + strings.ToLower(key)] = headerValues
	}

	for _, value := range request.Cookies() {
		params["cookie:" + value.Name] = value.Value
	}

	if interrogator.configCdn != nil && interrogator.configCdn.URL != "" {
		params["cdn:url"] = interrogator.configCdn.URL
	}

	params["url:href"] = pageUrl

	encodedParams := make(map[string]string)
	for key, value := range params {
		encodedValue, _ := util.EncodeUrl(value)
		encodedParams[key + ":encoded"] = encodedValue
	}

	for key, value := range encodedParams {
		params[key] = value
	}

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

func (interrogator *RequestInterrogator) interrogatePath(path string) map[string]string {
	returnParams := make(map[string]string)
	for _,url := range interrogator.configParams.Urls {
		r, _ := regexp.Compile(url.Pattern)
		matches := r.FindAllStringSubmatch(path, -1)
		if len(matches) > 0 {
			for i, match := range matches[0][1:] {
				returnParams[url.Names[i]] = match
			}
		}
	}
	return returnParams
}

func (interrogator *RequestInterrogator) interrogateParams(params url.Values) map[string]string {
	returnParams := make(map[string]string)
	if interrogator.configParams.Query != nil {
		for _, item := range interrogator.configParams.Query {
			queryItem := params.Get(item.Key)
			if queryItem != "" {
				returnParams[item.MapTo] = queryItem
			}
		}
	}
	return returnParams
}
