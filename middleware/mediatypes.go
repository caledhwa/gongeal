package middleware

import (
	"log"
	"net/http"
	"../config"
)

type RejectUnsupportedMediaTypeMiddleware struct {
	config *config.Config
}

func NewRejectUnsupportedMediaTypeMiddleware(config *config.Config) *RejectUnsupportedMediaTypeMiddleware {
	m := &RejectUnsupportedMediaTypeMiddleware{}
	m.config = config
	return m
}

func (middleware *RejectUnsupportedMediaTypeMiddleware) Handle(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("media types called")
		h.ServeHTTP(w,r)
	})
}

