package middleware

import (
	"log"
	"net/http"
	"github.com/gorilla/context"
	"github.com/caledhwa/gongeal/config"
	"github.com/caledhwa/gongeal/requestinterrogator"
	"github.com/caledhwa/gongeal/util"
)

type InterrogatorMiddleware struct {
	interrogator *requestinterrogator.RequestInterrogator
}

func NewInterrogatorMiddleware(config *config.Config) *InterrogatorMiddleware {
	m := &InterrogatorMiddleware{}
	m.interrogator = requestinterrogator.NewRequestInterrogator(&config.Parameters,&config.Cdn)
	return m
}

func (middleware *InterrogatorMiddleware) Handle(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		templateParameters := middleware.interrogator.InterrogateRequest(r)
		context.Set(r, "templateParameters", templateParameters)
		util.PrintJson(templateParameters)
		log.Println("interrogator called")
		h.ServeHTTP(w,r)
	})
}
