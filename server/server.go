package server

import (
	"github.com/caledhwa/gongeal/config"
	"github.com/caledhwa/gongeal/middleware"
	"net/http"
	"github.com/gorilla/context"
	"github.com/justinas/alice"
	"strconv"
)

func Start(port int, configuration *config.Config) {

	portString := ":" + strconv.Itoa(port)
	dropFavicon := middleware.NewFaviconMiddleware(configuration)

	// cache
	interrogateRequest := middleware.NewInterrogatorMiddleware(configuration)
	selectBackend := middleware.NewSelectBackendMiddleware(configuration)
	rejectUnsupportedMediaTypes := middleware.NewRejectUnsupportedMediaTypeMiddleware(configuration)
	passthrough := middleware.NewPassthroughMiddleware(configuration)

	// cookieParser
	backendProxy := middleware.NewBackendProxyMiddleware(configuration)

	chain := alice.New(dropFavicon.Handle,
		interrogateRequest.Handle,
		selectBackend.Handle,
		rejectUnsupportedMediaTypes.Handle,
		passthrough.Handle,
		backendProxy.Handle).ThenFunc(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		target := context.Get(r, "renderedTarget")
		if (target != nil) {
			renderedTarget := target.(string)
			w.Write([]byte(renderedTarget))
		}
	}))

	http.ListenAndServe(portString, chain)
}