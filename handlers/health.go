package handlers

import (
	"encoding/json"
	"net/http"
)

type HealthResponse struct {
	status string `json:"status"`
}

func (h *Handler) HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(HealthResponse{status: "ok"})
}
