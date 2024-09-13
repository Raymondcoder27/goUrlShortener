package main

import (
	"net/http"
	"sync"

	"github.com/go-chi/chi"
)

type Mapper struct {
	Mapping map[string]string
	Lock    sync.Mutex
}

var urlMapper Mapper

func init() {
	urlMapper = Mapper{
		Mapping: make(map[string]string),
	}
}

func main() {
	// Code
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server is running...."))
	})
	http.ListenAndServe(":3000", r)
}
