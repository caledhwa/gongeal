package common

import (
	"log"
	"net"
	"net/http"
)

func StartStubServer (port string) {

	log.Printf("Starting Stub at Port%s", port)

	// Serves static pages
	log.Println("Serving / - serves Stub html files for testing")
	server := &http.Server { Handler: &StaticHandler{} }

	log.Println("Listening on" + port)
	listener, err := net.Listen("tcp", port)
	if nil != err {
		log.Fatalln(err)
	}
	if err := server.Serve(listener); nil != err {
		log.Fatalln(err)
	}
}
