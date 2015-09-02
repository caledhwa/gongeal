package middleware

import (
	"log"
	"net/http"
	"regexp"
	"github.com/caledhwa/gongeal/config"
	"github.com/caledhwa/gongeal/util"
	"github.com/gorilla/context"
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

		if middleware.config.Backend != nil {
			var capturedBackend *config.Backend
			for _,backend := range middleware.config.Backend {
				if backend.Pattern != "" {
					backendMatched, _ :=  regexp.MatchString(backend.Pattern,r.URL.String())
					if (backendMatched) {
						log.Println("Found backend based on pattern")
						capturedBackend = &backend
						break
					}
				} else if backend.Fn != "" && middleware.config.SelectorFunctions[backend.Fn] != nil {
					log.Println("Found the function")
					rv := context.Get(r, "templateParameters")
					if middleware.config.SelectorFunctions[backend.Fn](r,rv.(map[string]string)) {
						capturedBackend = &backend
						break
					}
				}
			}

			if capturedBackend == nil {
				log.Println("Backend not found") //Log this as a warning
				w.WriteHeader(http.StatusNotFound)
			} else {
				log.Println("Backend selected")
				setBackendDefaults(capturedBackend)
				util.PrintJson(capturedBackend)

				// Render the target
				// Render the cache key
			}
		}

		h.ServeHTTP(w,r)
	})
}

func setBackendDefaults(capturedBackend *config.Backend) {

	truep := true
	falsep := false
	if capturedBackend.DontPassUrl == nil {
		capturedBackend.DontPassUrl = &truep
	}
	if capturedBackend.LeaveContentOnFail == nil {
		capturedBackend.LeaveContentOnFail = &truep
	}
	if capturedBackend.QuietFailure == nil {
		capturedBackend.QuietFailure = &falsep
	}
	if capturedBackend.ReplaceOuter == nil {
		capturedBackend.ReplaceOuter = &falsep
	}
}