package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	// create a new handler
	mux := http.NewServeMux()

	registerHandler(mux)

	err := http.ListenAndServe(":3000", mux)
	fmt.Println(err)
}

// register the handler function with handler
func registerHandler(mux *http.ServeMux) {
	mux.HandleFunc("/health", healthhandler)
	mux.HandleFunc("/", apiHandler)
}

// handler functions
func apiHandler(w http.ResponseWriter, r *http.Request) {
	reqDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Request", string(reqDump))
	fmt.Fprintf(w, "Hello World")
}

func healthhandler(w http.ResponseWriter, r *http.Request) {
	reqDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Request", string(reqDump))

	fmt.Fprintf(w, "OK, Working")
}
