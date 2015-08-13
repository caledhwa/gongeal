package middleware

import (
	"log"
	"net/http"
	"../config"
)

type SelectBackendMiddleware struct {
	 config *config.Config
}

func NewSelectBackendMiddleware(config *config.Config) *SelectBackendMiddleware {
	m := &SelectBackendMiddleware{}
	m.config = config
	return m
}

func (middleware *SelectBackendMiddleware) Handle(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("select backend called")
		h.ServeHTTP(w,r)
	})
}