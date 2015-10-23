package util

import (
	"net/url"
	"encoding/json"
	"log"
	"fmt"
	"regexp"
	"strconv"
	"net/http"
	"io/ioutil"
)

func EncodeUrl(str string) (string, error) {
	u, err := url.Parse(str)
	if err != nil {
		return "", err
	}
	return u.String(), nil
}

func PrintJson(x interface{}) {
	y,_ := json.Marshal(&x)
	log.Println(string(y))
}

func TimeToMillis(timeString string) int {

	r, _ := regexp.Compile("(\\d+)(.*)")
	matched := r.FindAllStringSubmatch(timeString,-1)
	num, _ := strconv.Atoi(matched[0][1])
	period := matched[0][2]
	value := 0

	switch period {
	case "s":
		value = num * 1000
	case "m":
		value = num * 1000*60
	case "h":
		value = num * 1000*60*60
	case "d":
		value = num * 1000*60*60*24
	default:
		value = num
	}
	return value
}

func CacheKeyToStatsD() {

}

func UrlToCacheKey() {

}

func UpdateTemplateVariables() {

}

func FilterCookies() {

}

func ReflectHeaderAndBody(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	jsonHeader, _ := json.Marshal(r.Header)
	fmt.Fprintf(w, "%s Data: %s<br/><pre>%s</pre>", r.Method, string(body), string(jsonHeader))
}

func LogRequest(r *http.Request) {
	log.Printf("[%s] %q \n", r.Method, r.URL.String())
}

func WriteHtmlOk(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
}