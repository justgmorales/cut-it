package handlers

import (
	"fmt"
	"net/http"
)

func (h *Handler) RedirectHandler(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")

	originalUrl, ok := h.Store.Get(slug)
	if !ok {
		http.Error(w, fmt.Sprintf("Slug %s not found", slug), http.StatusNotFound)
		return
	}

	// sets Location header and sends 307
	http.Redirect(w, r, originalUrl, http.StatusTemporaryRedirect)
}
