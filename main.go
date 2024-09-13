package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/go-chi/chi"
	"github.com/lithammer/shortuuid/v4"
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

	//insert
	insertMapping(key, u)

	log.Println("url mapped")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("http:localhost:3000/short/%s", key)))
}

func insertMapping(key string, u string) {
	urlMapper.Lock.Lock()
	defer urlMapper.Lock.Unlock()

	urlMapper.Mapping[key] = u
}

func redirectHandler(w http.ResponseWriter, r http.Request) {
	key := chi.URLParam(r, "key")
	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("url field is empty."))
	}
}

func main() {
	// Code
	r := chi.NewRouter()
	r.Post("/short-it", createShortURLHandler)
	r.Get("/short/{key}", redirectHandler)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server is running...."))
	})
	http.ListenAndServe(":3000", r)
}
