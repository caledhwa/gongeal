package common

import (
	"log"
	"net"
	"time"
	"net/http"
	"github.com/drone/routes"
)

func StartStubServer (port string) {

	log.Printf("Starting Stub at Port%s", port)

	mux := routes.New()

	// Serves static pages
	log.Println("Serving / - serves Stub html files for testing")
	mux.Get("/:param", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Filepath:%s\n","../common/" + r.URL.Path[1:])
		t1 := time.Now()
		http.ServeFile(w, r, "../common/" + r.URL.Path[1:])
		t2 := time.Now()
		log.Printf("[STUB] [%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
	})

	log.Println("Listening on" + port)
	server := &http.Server{Handler: mux}
	listener, err := net.Listen("tcp", port)
	if nil != err {
		log.Fatalln(err)
	}
	if err := server.Serve(listener); nil != err {
		log.Fatalln(err)
	}
}
