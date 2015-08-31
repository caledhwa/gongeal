package middleware

import (
	"log"
	"net/http"
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

					// if the pattern exists, return bool after running pattern vs. URL
				} else if backend.Fn != "" && middleware.config.SelectorFunctions[backend.Fn] != nil {
					log.Println("Found the function")
					rv := context.Get(r, "templateParameters")
					if middleware.config.SelectorFunctions[backend.Fn](r,rv.(map[string]string)) {
						capturedBackend = &backend
						break;
					}
				}
			}

			if capturedBackend == nil {
				log.Println("Backend not found") //Log this as a warning
				w.WriteHeader(http.StatusNotFound)
			} else {
				log.Println("Backend selected")
				util.PrintJson(capturedBackend)
				// Add defaults
				// Render the target
				// Render the cache key
			}
		}

		h.ServeHTTP(w,r)
	})
}