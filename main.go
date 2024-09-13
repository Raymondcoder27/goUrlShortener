package main

import "net/http"

func main() {
	// Code
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server is running...."))
	})

}
