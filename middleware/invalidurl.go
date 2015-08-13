package middleware

import (
	"log"
	"net/http"
	"../config"
)

type CleanInvalidUriMiddleware struct {
	config *config.Config
}

func NewCleanInvalidUriMiddleware(config *config.Config) *CleanInvalidUriMiddleware {
	m := &CleanInvalidUriMiddleware{}
	m.config = config
	return m
}

func (middleware *CleanInvalidUriMiddleware) Handle(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("clean invalid uri called")
		h.ServeHTTP(w,r)
	})
}
