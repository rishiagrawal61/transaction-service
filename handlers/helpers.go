package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, error error) {
	log.Printf("Application Error: %s", error)
	writeJSON(w, status, map[string]string{"error": error.Error()})
}

// writeValidationErrors returns a 422 with a map of field → error message.
func writeValidationErrors(w http.ResponseWriter, errs map[string]string) {
	log.Printf("[validation] failed: %v", errs)
	writeJSON(w, http.StatusUnprocessableEntity, map[string]any{
		"error":  "validation failed",
		"fields": errs,
	})
}
