package main

import (
	"fmt"
	"net/http"

	"rsc.io/quote/v4"
)

func renderHomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you've requested %s\n", r.URL.Path)
	fmt.Fprintf(w, "Fun quote: \t\"%s\" \n", quote.Go())

}

func startServer() {
	http.HandleFunc("/", renderHomePage)

	http.ListenAndServe(":80", nil)
}

func main() {
	startServer()
}
