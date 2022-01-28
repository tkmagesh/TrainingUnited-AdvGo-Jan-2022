package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			w.Write([]byte("All the products will be returned"))
		case "POST":
			w.Write([]byte("A new product will be created"))
		case "PUT":
			w.Write([]byte("An existing product will be updated"))
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
	http.ListenAndServe(":8080", nil)
}
