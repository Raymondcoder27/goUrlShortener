package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	// Code
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server is running...."))
	})
	http.ListenAndServe(":3000", r)
}
