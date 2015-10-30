package main

import (
	"github.com/drone/routes"
	"github.com/caledhwa/gongeal/util"
	"net/http"
	"log"
	"strings"
	"fmt"
	"time"
	"net"
	"strconv"
)

func main() {
	StubServer(9000)
}

func StubServer(port int) {

	portString := ":" + strconv.Itoa(port)

	mux := routes.New()

	mux.Get("/broken", func(w http.ResponseWriter, r *http.Request) {
		// The Routes library doesn't allow hijacking
		// With hijacking you can close the connection with no response
		// To simulate a close with empty response, a panic will end the response empty
		panic("Cannot close connection with empty response using github.com/drone/routes")
	})

	mux.Get("/faulty", faultyFunction)
	mux.Get("/cb-faulty", faultyFunction)
	mux.Get("/cb-faulty-default", faultyFunction)
	mux.Get("/teaching-resource/:resourceStub", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "complex.html")
		util.LogRequest(r)
	})

	mux.Get("/nocache", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		w.Header().Set("cache-control", "private, max-age=0, no-cache")
		util.WriteHtmlOk(w)
		fmt.Fprintf(w, "%d-%02d-%02dT%02d:%02d:%02d-00:00", t.Year(), t.Month(), t.Day(),
			t.Hour(), t.Minute(), t.Second())
		util.LogRequest(r)
	})

	mux.Get("/maxage", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		w.Header().Set("cache-control", "private, max-age=5000")
		util.WriteHtmlOk(w)
		fmt.Fprintf(w, "%d-%02d-%02dT%02d:%02d:%02d-00:00", t.Year(), t.Month(), t.Day(),
			t.Hour(), t.Minute(), t.Second())
		util.LogRequest(r)
	})

	mux.Get("/302", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusFound)
		fmt.Fprint(w, "")
		util.LogRequest(r)
	})

	mux.Get("/set-cookie", func(w http.ResponseWriter, r *http.Request) {
		if !(strings.Index(r.URL.String(), "faulty=true") >= 0) {
			w.Header().Set("Content-Type", "text/html")
			w.Header().Set("Set-Cookie", "test=bob")
			util.WriteHtmlOk(w)
			fmt.Fprint(w, "OK")
		} else {
			w.Header().Set("Content-Type", "text/html")
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "Faulty service broken")
		}
		util.LogRequest(r)
	})

	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		util.WriteHtmlOk(w)
		fmt.Fprint(w, "OK-Root")
		util.LogRequest(r)
	})

	log.Println("Listening on " + portString)
	server := &http.Server{Handler: mux}
	listener, err := net.Listen("tcp", portString)
	if nil != err {
		log.Fatalln(err)
	}
	if err := server.Serve(listener); nil != err {
		log.Fatalln(err)
	}
}

func faultyFunction(w http.ResponseWriter, r *http.Request) {
	time.Sleep(100 * time.Millisecond)
	if !(strings.Index(r.URL.String(), "faulty=true") >= 0) {
		w.Header().Set("Content-Type", "text/html")
		util.WriteHtmlOk(w)
		fmt.Fprint(w, "Faulty service managed to serve good content!")
	} else {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Faulty service broken")
	}
	util.LogRequest(r)
}

