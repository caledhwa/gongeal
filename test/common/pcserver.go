package common

import (
	"log"
	"net"
	"net/http"
	"strconv"
)

func StartPageCompositionServer (port int, hostname string, eventHandler func(), config string ) {

	var configValue string
	if config != "" {
		configValue = config
	} else {
		configValue = "testConfig"
	}
	configValue += ".json"

	//var config = require('./' + (configFile || 'testConfig') + '.json');

	portString := ":" + strconv.Itoa(port)
	log.Printf("Starting PcServer at Port: %v\n", portString)

	// Serves static pages
	log.Println("Serving / - serves Page Composition html files for testing")


	log.Println("Listening on " + portString)

	server := &http.Server { Handler: &StaticHandler{} }

	listener, err := net.Listen("tcp", portString)
	if nil != err {
		log.Fatalln(err)
	}
	if err := server.Serve(listener); nil != err {
		log.Fatalln(err)
	}
}


