package middleware

import (
	"log"
	"net/http"
	"github.com/caledhwa/gongeal/config"
	"github.com/caledhwa/gongeal/util"
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

		if middleware.config.Backend == nil {
			log.Println("Backend is empty")
		} else {
			util.PrintJson(middleware.config.Backend)
		}


		h.ServeHTTP(w,r)
	})
}