package main

import (
	"net/http"
	"github.com/arulthayalan/url-shortener/urlshort"
	"fmt"
	"log"
)

const PORT int = 8080

func main() {
	mux := defaultMux()
	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://pkg.go.dev/gopkg.in/yaml.v2",
	}

	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`
	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Starting the server on :%d\n", PORT)

	//http.HandleFunc("/", defaultHandler)
	
	//default hanlder
	//log.Fatal(http.ListenAndServe(host, mux))

	//Map handler
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d",PORT), yamlHandler))

}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", defaultHandler)
	return mux
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to GOLang server side programming")
} 