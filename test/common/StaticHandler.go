package common

import (
	"net/http"
	"log"
	"time"
)

type StaticHandler struct {}

func (self *StaticHandler) ServeHTTP (w http.ResponseWriter, r *http.Request) {
	t1 := time.Now()
	http.ServeFile(w, r, "../common" + r.URL.Path[0:])
	t2 := time.Now()
	log.Printf("[PAGE COMPOSITION] [%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
}
