package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

var cnt int

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Get("/hello", Hello)
	httpPort := ":"
	if env, ok := os.LookupEnv("PORT"); !ok {
		httpPort += "8080"
	} else {
		httpPort += env
	}
	log.Panic(http.ListenAndServe(httpPort, r))
}

func Hello(w http.ResponseWriter, r *http.Request) {
	cnt++
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, guy! You are %d", cnt)
}
