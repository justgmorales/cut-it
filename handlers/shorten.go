package handlers

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type incomingURL struct {
	OriginalURL string `json:"original_url"`
}

type outboundURL struct {
	ShortenedURL string `json:"shortened_url"`
}

func (h *Handler) ShortenHandler(w http.ResponseWriter, r *http.Request) {
	// never runs? does mux prevent invalid requests from coming through to handler?
	// if r.Method != http.MethodPost {
	// 	http.Error(w, "Invalid HTTP method for endpoint /shorten", http.StatusMethodNotAllowed)
	// 	return
	// }

	var incomingUrl incomingURL

	err := json.NewDecoder(r.Body).Decode(&incomingUrl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest) // status: 400
		return
	}

	u, err := url.Parse(incomingUrl.OriginalURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity) // status: 422
		return
	}

	if u.Scheme != "http" && u.Scheme != "https" {
		http.Error(w, "Invalid scheme in URL", http.StatusUnprocessableEntity)
		return
	} else if u.Host == "" || u.Host == h.BaseURL.Host {
		http.Error(w, "Invalid host in URL", http.StatusUnprocessableEntity)
		return
	}

	slug, err := h.Store.Set(u.String())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // status :500
		return
	}

	shortenedURL, err := url.JoinPath(h.BaseURL.String(), slug)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	payload := outboundURL{
		ShortenedURL: shortenedURL,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(payload)

	// do we need this? writeheader locks in the status code so this wont even work
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
