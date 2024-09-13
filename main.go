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

func createShortURLHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	u := r.Form.Get("URL")
	if u == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("url field is empty."))
	}

	//generate key
	key := shortuuid.New()
}

func insertMapping(key string, u string) {
	urlMapper.Lock.Lock()
	defer urlMapper.Lock.Unlock()

	urlMapper.Mapping[key] = u
}

func main() {
	// Code
	r := chi.NewRouter()
	r.Post("/short-it", createShortURL)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server is running...."))
	})
	http.ListenAndServe(":3000", r)
}
