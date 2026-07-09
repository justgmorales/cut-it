package handlers

import (
	"net/http"
	"net/url"

	"github.com/justgmorales/cut-it/store"
)

type Handler struct {
	Store   *store.Store
	BaseURL *url.URL
}

func RegisterRoutes(mux *http.ServeMux, h *Handler) {
	mux.HandleFunc("POST /shorten", h.ShortenHandler)
	mux.HandleFunc("GET /{slug}", h.RedirectHandler)
}
