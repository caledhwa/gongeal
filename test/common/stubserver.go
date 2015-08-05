package common

import (
	"log"
	"net"
	"net/http"
	"strconv"
)

func StartStubServer (port int) {

	portString := ":" + strconv.Itoa(port)
	log.Printf("Starting Stub at Port: %v\n", portString)

	// Serves static pages
	log.Println("Serving / - serves Stub html files for testing")
	server := &http.Server { Handler: &StaticHandler{} }

	log.Println("Listening on " + portString)
	listener, err := net.Listen("tcp", portString)
	if nil != err {
		log.Fatalln(err)
	}
	if err := server.Serve(listener); nil != err {
		log.Fatalln(err)
	}
}
