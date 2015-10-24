package main

import (
	"encoding/json"
	"github.com/caledhwa/gongeal/server"
	"github.com/caledhwa/gongeal/config"
	"log"
	"net/http"
	"os"
)

func main() {

	configFile, _ := os.Open("example/config.json")
	jsonParser := json.NewDecoder(configFile)
	var configuration config.Config
	_ = jsonParser.Decode(&configuration)

	configuration.SelectorFunctions = make(map[string]config.BackendSelectorFunction)
	configuration.SelectorFunctions["selectGoogle"] = selectGoogle

	server.Start(8000, &configuration)
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
