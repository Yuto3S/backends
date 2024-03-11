package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func createNewMiddleware() Middleware {
	middleware := func(next http.HandlerFunc) http.HandlerFunc {
		handler := func(writer http.ResponseWriter, req *http.Request) {
			next(writer, req)
		}

		return handler
	}
	return middleware
}

func Logging() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, req *http.Request) {
			start := time.Now()
			defer func() {
				log.Println(req.URL.Path, time.Since(start))
			}()

			f(writer, req)
		}
	}
}

func Method(method string) Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, req *http.Request) {
			if req.Method != method {
				http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}

			f(writer, req)
		}
	}
}

func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, middleware := range middlewares {
		f = middleware(f)
	}
	return f
}

func HelloWorld(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello World :) This should log the time and return an error if not a GET")
}
