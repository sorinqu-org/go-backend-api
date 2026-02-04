package json

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func Write(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func Read(r *http.Request, w http.ResponseWriter, v any) error {
	err := json.NewDecoder(r.Body).Decode(v)
	if err != nil {
		slog.Error(
			"Failed to Decode data",
			"error",
			err.Error(),
			"data",
			r.Body,
		)
		http.Error(w, "Invalid JSON model", http.StatusNotAcceptable)
	}
	return err
}
