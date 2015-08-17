package main

import (
	"net/http"
	"github.com/justinas/alice"
	"github.com/caledhwa/gongeal/middleware"
	"github.com/caledhwa/gongeal/config"
)

func main() {

	config := &config.Config{}

	cleanInvalidUri := middleware.NewCleanInvalidUriMiddleware(config)
	dropFavicon := middleware.NewFaviconMiddleware(config)
	// cache
	interrogateRequest := middleware.NewInterrogatorMiddleware(config)
	selectBackend := middleware.NewSelectBackendMiddleware(config)
	rejectUnsupportedMediaTypes := middleware.NewRejectUnsupportedMediaTypeMiddleware(config)
	passthrough := middleware.NewPassthroughMiddleware(config)
	// cookieParser
	backendProxy := middleware.NewBackendProxyMiddleware(config)

	chain := alice.New(	cleanInvalidUri.Handle,
					   	dropFavicon.Handle,
					   	interrogateRequest.Handle,
						selectBackend.Handle,
						rejectUnsupportedMediaTypes.Handle,
						passthrough.Handle,
						backendProxy.Handle).ThenFunc(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world!"))
	}))

	http.ListenAndServe(":8000", chain)
}
