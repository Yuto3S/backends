package main

import (
	"fmt"
	"net/http"
	"rsc.io/quote/v4"
)

func startServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { 
		fmt.Fprintf(w, "Hello, you've requested %s\n", r.URL.Path)
	})

	http.ListenAndServe(":80", nil)
}

func main() {
	fmt.Println(quote.Go())
	startServer()
}