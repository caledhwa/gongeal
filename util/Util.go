package util

import (
	"net/url"
	"encoding/json"
	"log"
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