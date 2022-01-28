package main

import (
	"fmt"
	"net/http"
	"time"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func foo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("foo"))
}

func bar(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bar"))
}

/* middleware */
func logger(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Before request", r.URL.Path)
		handler(w, r)
		fmt.Println("After request", r.URL.Path)
	}
}

func profile(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer func() {
			end := time.Now()
			elapsed := end.Sub(start)
			fmt.Println("Request took", elapsed)
		}()
		handler(w, r)
	}
}

func chain(handler http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		handler = m(handler)
	}
	return handler
}

func main() {
	/*
		http.HandleFunc("/foo", profile(logger(foo)))
		http.HandleFunc("/bar", profile(logger(bar)))
	*/
	http.HandleFunc("/foo", chain(foo, profile, logger))
	http.HandleFunc("/bar", chain(bar, profile, logger))

	http.ListenAndServe(":8080", nil)
}
