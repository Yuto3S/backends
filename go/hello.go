package main

import (
	"fmt"
	"net/http"

	"rsc.io/quote/v4"
)

func renderHomePage(writer http.ResponseWriter, request *http.Request) {
	var test_int_for_debug_dlv int = 5
	test_debug := 4
	fmt.Printf("Test sum: %d + %d = %d\n", test_int_for_debug_dlv, test_debug, test_debug+test_int_for_debug_dlv)

	// Accessing http://localhost/kekz/eze?param=testeze will print "testeze"
	fmt.Printf("%s", request.URL.Query().Get("param"))

	fmt.Fprintf(writer, "Hello, you've requested %s\n", request.URL.Path)
	fmt.Fprintf(writer, "Fun quote: \t\"%s\" \n", quote.Go())

}

func startServer() {
	http.HandleFunc("/", renderHomePage)

	http.ListenAndServe(":80", nil)
}

func main() {
	startServer()
}
