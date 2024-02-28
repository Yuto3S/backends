package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func defineBookRoutes(router *mux.Router) {
	bookrouter := router.PathPrefix("/books").Subrouter()
	bookrouter.HandleFunc("/", allBooks)

	// http://localhost/books/MyAmazingBook/page/1
	bookrouter.HandleFunc("/{title}/page/{page}", renderBookPage)

	// $ curl http://localhost/books/backendgobook
	// Book exists: backendgobook
	// $ curl http://localhost/books/backendgobook -X DELETE
	// Book deleted: backendgobook
	// $ curl http://localhost/books/backendgobook -X POST
	// Book created: backendgobook

	bookrouter.HandleFunc(bookTitleHTTPPath, createBook).Methods("POST")
	// localhost doesn't handle https for now
	bookrouter.HandleFunc(bookTitleHTTPPath, getBook).Methods("GET").Host("localhost").Schemes("http")
	bookrouter.HandleFunc(bookTitleHTTPPath, deleteBook).Methods("DELETE")
}

func allBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Here is the list of all the available books: ... [work in progress] ...")
}

func renderBookPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "Book requested: %s on page %s", vars["title"], vars["page"])
}

func createBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "Book created: %s\n", vars["title"])
}

func getBook(w http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	fmt.Fprintf(w, "Book exists: %s\n", vars["title"])
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "Book deleted: %s\n", vars["title"])
}

// func bookLocalHost(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Local Host")
// }

func startServerWithRouter() {
	router := mux.NewRouter()
	defineBookRoutes(router)
	http.ListenAndServe(":80", router)
}
