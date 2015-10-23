package main

import (
	"fmt"
	"net"
	"time"
	"net/http"
	"log"
	"math/rand"
	"encoding/json"
	"io/ioutil"
	"github.com/drone/routes"
)

var port = ":5001"

func main() {

	log.Println("Gongeal Test Server: A Go Port of Compoxure for Composition of UX")
	log.Println("Setting up handlers... on port", port)


	mux := routes.New()

	// Serves static pages
	log.Println("Serving /static - serves html files for testing")
	mux.Get("/:param", func(w http.ResponseWriter, r *http.Request) {
		path := "static/" + r.URL.Path[1:]
		log.Printf("Filepath:%s\n",path)
		t1 := time.Now()
		http.ServeFile(w, r, path)
		t2 := time.Now()
		log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
	})

	// Example of a dynamic content for testing
	log.Println("Serving /dynamic - Shows dynamic time content")
	mux.Get("/dynamic", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		w.Header().Set("Content-Type", "text/html")
		w.Header().Set("x-static|service|top", "100")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "This is some dynamic comment: %d-%02d-%02dT%02d:%02d:%02d-00:00", t.Year(), t.Month(), t.Day(),
			t.Hour(), t.Minute(), t.Second())
		logRequest(r)
	})

	// 500
	log.Println("Serving /500 - Simulates a server error")
	mux.Get("/500", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "This is an error.")
		logRequest(r)
	})

	// 403
	fmt.Println("Serving /403 - Simulates a Forbidden")
	mux.Get("/403", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprint(w, "Unauthorised error.")
		logRequest(r)
	})

	// Broken
	log.Println("Serving /broken - Rudely ends the request (Serves empty request)")
	mux.Get("/broken", func(w http.ResponseWriter, r *http.Request) {
		hj, _ := w.(http.Hijacker)
		conn, _, err := hj.Hijack()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer conn.Close()
		logRequest(r)
	})

	// 500
	log.Println("Serving /faulty - randomly returns a 200 or 500")
	mux.Get("/faulty", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(100 * time.Millisecond)
		if rand.Float32() > 0.5 {
			writeHtmlOk(w)
			fmt.Fprint(w, "Faulty service managed to serve good content!")
		} else {
			w.Header().Set("Content-Type", "text/html")
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "Faulty Service Broken")
		}
		logRequest(r)
	})

	fmt.Println("Serving /slow - Returns a delayed response (slow service)")
	mux.Get("/slow", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(200 * time.Millisecond)
		writeHtmlOk(w)
		fmt.Fprint(w, "This is a slow service.")
		logRequest(r)
	})

	fmt.Println("Serving /post - Reflects a POST request")
	mux.Post("/post", func(w http.ResponseWriter, r *http.Request) {
		writeHtmlOk(w)
		reflectHeaderAndBody(w, r)
		logRequest(r)
	})

	fmt.Println("Serving /put - Reflects a PUT request")
	mux.Put("/put", func(w http.ResponseWriter, r *http.Request) {
		writeHtmlOk(w)
		reflectHeaderAndBody(w, r)
		logRequest(r)
	})

	fmt.Println("Serving /cdn - Simulates a CDN (Format:/cdn/:environment/:version/html/:file)")
	mux.Get("/cdn/:environment/:version/html/:file", func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Environment: %s, Version: %s, File: %s", params.Get(":environment"), params.Get(":version"), params.Get(":file"))
		logRequest(r)
	})

	fmt.Println("Listening on" + port)
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

func reflectHeaderAndBody(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	jsonHeader, _ := json.Marshal(r.Header)
	fmt.Fprintf(w, "%s Data: %s<br/><pre>%s</pre>", r.Method, string(body), string(jsonHeader))
}

func logRequest(r *http.Request) {
	log.Printf("[%s] %q \n", r.Method, r.URL.String())
}

func writeHtmlOk(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
}
