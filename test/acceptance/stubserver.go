package acceptance

import (
	"log"
	"net"
	"net/http"
	"strconv"
	"github.com/drone/routes"
	"time"
	"fmt"
	"github.com/caledhwa/gongeal/util"
	"math/rand"
)

func StartStubServer (port int) {

	portString := ":" + strconv.Itoa(port)
	log.Printf("Starting Stub at Port: %v\n", portString)

	mux:= routes.New()

	mux.Get("/404backend", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../common/test404.html")
		util.LogRequest(r)
	})

	mux.Get("/500", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "500.")
		util.LogRequest(r)
	})

	mux.Get("/403", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprint(w, "Unauthorised error.")
		util.LogRequest(r)
	})

	mux.Get("/404", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "404")
		util.LogRequest(r)
	})

	mux.Get("/delayed", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(100 * time.Millisecond)
		util.WriteHtmlOk(w)
		fmt.Fprint(w, "Delayed by 100ms")
		util.LogRequest(r)
	})

	mux.Get("/timeout", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(6000 * time.Millisecond)
		util.WriteHtmlOk(w)
		fmt.Fprint(w, "Delayed by 6seconds")
		util.LogRequest(r)
	})

	mux.Get("/broken", func(w http.ResponseWriter, r *http.Request) {
		// The Routes library doesn't allow hijacking
		// With hijacking you can close the connection with no response
		// To simulate a close with empty response, a panic will end the response empty
		panic("Cannot close connection with empty response using github.com/drone/routes")
	})

	mux.Get("/faulty", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(100 * time.Millisecond)
		if rand.Float32() > 0.5 {
			util.WriteHtmlOk(w)
			fmt.Fprint(w, "Faulty service managed to serve good content!")
		} else {
			w.Header().Set("Content-Type", "text/html")
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "Faulty Service Broken")
		}
		util.LogRequest(r)
	})

	mux.Get("/403backend", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "test403.html")
		util.LogRequest(r)
	})

	mux.Get("/302backend", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "test302.html")
		util.LogRequest(r)
	})

	mux.Get("/ignore404backend", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "ignore404.html")
		util.LogRequest(r)
	})

	mux.Get("/selectFnBackend", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "test403.html")
		util.LogRequest(r)
	})

	mux.Get("/noCacheBackend", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "noCacheBackend.html")
		util.LogRequest(r)
	})

	mux.Get("/:param", func(w http.ResponseWriter, r *http.Request) {
		path := "../common/" + r.URL.Path[1:]
		log.Printf("Filepath:%s\n",path)
		t1 := time.Now()
		http.ServeFile(w, r, path)
		t2 := time.Now()
		log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
	})

	server := &http.Server{Handler: mux}
	listener, err := net.Listen("tcp", portString)
	if nil != err {
		log.Fatalln(err)
	}
	if err := server.Serve(listener); nil != err {
		log.Fatalln(err)
	}
}
