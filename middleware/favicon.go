package middleware

import (
	"log"
	"net/http"
	"github.com/caledhwa/gongeal/config"
)

type FaviconMiddleware struct {
	config *config.Config
}

func NewFaviconMiddleware(config *config.Config) *FaviconMiddleware {
	m := &FaviconMiddleware{}
	m.config = config
	return m
}

func (middleware *FaviconMiddleware) Handle(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Favicon called")
		if r.URL.Path[1:] == "favicon.ico" {
			w.Header().Set("Content-Type", "image/x-icon")
			w.WriteHeader(http.StatusOK)
			log.Println("Dropped favicon request")
		}
		h.ServeHTTP(w,r)
	})
}
