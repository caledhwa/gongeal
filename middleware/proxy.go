package middleware

import (
	"log"
	"net/http"
	"../config"
)

type BackendProxyMiddleware struct {
	config *config.Config
}

func NewBackendProxyMiddleware(config *config.Config) *BackendProxyMiddleware {
	m := &BackendProxyMiddleware{}
	m.config = config
	return m
}

func (middleware *BackendProxyMiddleware) Handle(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("backend proxy called")
		h.ServeHTTP(w,r)
	})
}