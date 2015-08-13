package middleware

import (
	"log"
	"net/http"
	"../config"
)

type PassthroughMiddleware struct {
	config *config.Config
}

func NewPassthroughMiddleware(config *config.Config) *PassthroughMiddleware {
	m := &PassthroughMiddleware{}
	m.config = config
	return m
}

func (middleware *PassthroughMiddleware) Handle(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("passthrough called")
		h.ServeHTTP(w,r)
	})
}
