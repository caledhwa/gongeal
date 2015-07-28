package common

import (
	"log"
	"time"
	"net"
	"net/http"
	"github.com/drone/routes"
)

func StartPageCompositionServer (port string) {

	log.Printf("Starting PcServer at Port%s\n", port)

	mux := routes.New()

	// Serves static pages
	log.Println("Serving / - serves Page Composition html files for testing")
	mux.Get("/:param", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Filepath:%s\n","../common/" + r.URL.Path[1:])
		t1 := time.Now()
		http.ServeFile(w, r, "../common/" + r.URL.Path[1:])
		t2 := time.Now()
		log.Printf("[PAGE COMPOSITION] [%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
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
