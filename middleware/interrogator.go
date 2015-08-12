package middleware

import (
	"log"
	"net/http"
	"../config"
)

type InterrogatorMiddleware struct {
	config *config.Config
}

func NewInterrogatorMiddleware(config *config.Config) *InterrogatorMiddleware {
	m := &InterrogatorMiddleware{}
	m.config = config
	return m
}

func (middleware *InterrogatorMiddleware) Handle(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("interrogator called")
		h.ServeHTTP(w,r)
	})
}
