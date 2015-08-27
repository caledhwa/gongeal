package main

import (
	"net/http"
	"github.com/justinas/alice"
	"github.com/caledhwa/gongeal/middleware"
	"github.com/caledhwa/gongeal/config"
	"os"
	"encoding/json"
)

func main() {

	// TEMP CONFIG CODE TO PULL FROM CONFIG IN TEST FOLDER
	configFile, _ := os.Open("test/common/testConfig.json")
	jsonParser := json.NewDecoder(configFile)
	var configuration config.Config
	_ = jsonParser.Decode(&configuration)

	dropFavicon := middleware.NewFaviconMiddleware(&configuration)
	// cache
	interrogateRequest := middleware.NewInterrogatorMiddleware(&configuration)
	selectBackend := middleware.NewSelectBackendMiddleware(&configuration)
	rejectUnsupportedMediaTypes := middleware.NewRejectUnsupportedMediaTypeMiddleware(&configuration)
	passthrough := middleware.NewPassthroughMiddleware(&configuration)
	// cookieParser
	backendProxy := middleware.NewBackendProxyMiddleware(&configuration)

	chain := alice.New(	dropFavicon.Handle,
					   	interrogateRequest.Handle,
						selectBackend.Handle,
						rejectUnsupportedMediaTypes.Handle,
						passthrough.Handle,
						backendProxy.Handle).ThenFunc(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
						w.Write([]byte("Final handler - eventually backend Proxy"))
	}))

	http.ListenAndServe(":8000", chain)
}
