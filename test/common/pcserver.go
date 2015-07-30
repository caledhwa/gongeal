package common

import (
	"log"
	"net"
	"net/http"
)

func StartPageCompositionServer (port string) {

	log.Printf("Starting PcServer at Port%s\n", port)

	// Serves static pages
	log.Println("Serving / - serves Page Composition html files for testing")


	log.Println("Listening on" + port)

	server := &http.Server { Handler: &StaticHandler{} }

	listener, err := net.Listen("tcp", port)
	if nil != err {
		log.Fatalln(err)
	}
	if err := server.Serve(listener); nil != err {
		log.Fatalln(err)
	}
}


