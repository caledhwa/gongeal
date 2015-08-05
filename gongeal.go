package gongeal

import (
	"fmt"
	"net/http"
	"os"
	"github.com/hoisie/mustache"
	"golang.org/x/net/html"
)

func main() {

	port := ":1338"

	fmt.Println("Setting up handler...")
	http.HandleFunc("/", handle)
	http.HandleFunc("/mustache", mustacheTest)

	fmt.Println("Listening on" + port)
	http.ListenAndServe(port, nil)
}

func mustacheTest(w http.ResponseWriter, r *http.Request) {
	data := mustache.Render("hello {{c}}", map[string]string{"c": "world"})
	fmt.Fprintln(w, "Mustache Test: "+data)
}

func handle(w http.ResponseWriter, r *http.Request) {
	file, _ := os.Open("test.html")
	tokenizer := html.NewTokenizer(file)

	fmt.Fprintln(w, "----------- Start Tokens ---------------")

	for {
		if tokenizer.Next() == html.ErrorToken {
			break
		}
		fmt.Fprintln(w, "Token: "+tokenizer.Token().String())
	}

	fmt.Fprintln(w, "----------- End Tokens ---------------")
}
