package main

import (
	"encoding/json"
	"github.com/caledhwa/gongeal/config"
	"github.com/caledhwa/gongeal/middleware"
	"github.com/gorilla/context"
	"github.com/justinas/alice"
	"log"
	"net/http"
	"os"
)

func main() {

	// TEMP CONFIG CODE TO PULL FROM CONFIG IN TEST FOLDER
	//configFile, _ := os.Open("test/common/testConfig.json")
	configFile, _ := os.Open("example/config.json")
	jsonParser := json.NewDecoder(configFile)
	var configuration config.Config
	_ = jsonParser.Decode(&configuration)

	configuration.SelectorFunctions = make(map[string]config.BackendSelectorFunction)
	configuration.SelectorFunctions["selectGoogle"] = selectGoogle

	dropFavicon := middleware.NewFaviconMiddleware(&configuration)
	// cache
	interrogateRequest := middleware.NewInterrogatorMiddleware(&configuration)
	selectBackend := middleware.NewSelectBackendMiddleware(&configuration)
	rejectUnsupportedMediaTypes := middleware.NewRejectUnsupportedMediaTypeMiddleware(&configuration)
	passthrough := middleware.NewPassthroughMiddleware(&configuration)
	// cookieParser
	backendProxy := middleware.NewBackendProxyMiddleware(&configuration)

	chain := alice.New(dropFavicon.Handle,
		interrogateRequest.Handle,
		selectBackend.Handle,
		rejectUnsupportedMediaTypes.Handle,
		passthrough.Handle,
		backendProxy.Handle).ThenFunc(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		renderedTarget := context.Get(r, "renderedTarget").(string)
		w.Write([]byte(renderedTarget))
	}))

	http.ListenAndServe(":8000", chain)
}

func selectGoogle(r *http.Request, parameters map[string]string) bool {
	log.Println("Executing the selectGoogle function")
	if _, ok := parameters["query:google"]; ok {
		log.Println("Google Found.")
		return true
	} else {
		return false
	}
}
