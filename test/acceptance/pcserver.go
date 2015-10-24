package acceptance

import (
	"log"
	"strconv"
	"github.com/caledhwa/gongeal/server"
	"github.com/caledhwa/gongeal/config"
	"os"
	"encoding/json"
)

func StartPageCompositionServer (port int, hostname string, eventHandler func(), configFileName string) {

	if configFileName == "" {
		configFileName = "testConfig"
	}

	configFilePath := "../common/" + configFileName + ".json"

	configFile, _ := os.Open(configFilePath)

	jsonParser := json.NewDecoder(configFile)
	var configuration config.Config
	_ = jsonParser.Decode(&configuration)

	configuration.SelectorFunctions = make(map[string]config.BackendSelectorFunction)

	portString := ":" + strconv.Itoa(port)
	log.Printf("Starting Gongeal Server at Port: %v\n", portString)

	server.Start(port, &configuration)
}


