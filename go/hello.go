package main

import (
	"fmt"
	"html/template"
	"net/http"

	"rsc.io/quote/v4"
)

var bookTitleHTTPPath string = "/{title}"

func renderHomePage(writer http.ResponseWriter, request *http.Request) {
	// var test_int_for_debug_dlv int = 5
	// test_debug := 4
	// fmt.Printf("Test sum: %d + %d = %d\n", test_int_for_debug_dlv, test_debug, test_debug+test_int_for_debug_dlv)

	// Accessing http://localhost/kekz/eze?param=testeze will print "testeze"
	fmt.Printf("%s", request.URL.Query().Get("param"))

	fmt.Fprintf(writer, "Hello, you've requested %s\n", request.URL.Path)
	fmt.Fprintf(writer, "Fun quote: \t\"%s\" \n", quote.Go())

}

func handleRenderStatic() {
	// http://localhost/static/pepelaugh.gif
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
}

func startServer() {
	http.HandleFunc("/", renderHomePage)
	handleRenderStatic()

	http.ListenAndServe(":80", nil)
}

func startServerWithTemplate() {
	tmpl := template.Must(template.ParseFiles("example_template.gohtml", "example_template_list_of_tasks.html"))
	http.HandleFunc("/", func(writer http.ResponseWriter, r *http.Request) {
		renderTemplate(writer, tmpl)
	})
	http.ListenAndServe(":80", nil)
}

func startServerWithForm() {
	tmpl := template.Must(template.ParseFiles("form.html"))
	http.HandleFunc("/", func(writer http.ResponseWriter, req *http.Request) {
		handleSubmitDetails(writer, req, tmpl)
	})
	http.ListenAndServe(":80", nil)
}

func main() {
	// startServer()
	// startServerWithRouter()
	// initTables()
	// startServerWithTemplate()
	startServerWithForm()
}
