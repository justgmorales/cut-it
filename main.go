package main

import (
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/justgmorales/cut-it/handlers"
	"github.com/justgmorales/cut-it/store"
)

const baseURL = "http://localhost:8080"

func main() {
	// convert baseURL to a URL object
	u, err := url.Parse(baseURL)
	if err != nil {
		return
	}

	mux := http.NewServeMux()

	s := store.NewStore()
	h := &handlers.Handler{
		Store:   s,
		BaseURL: u,
	}

	handlers.RegisterRoutes(mux, h)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Second, // small payloads justify faster timeouts
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	// TODO: graceful shutdowns on err
	log.Fatal(srv.ListenAndServe())
}
