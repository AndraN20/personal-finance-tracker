package api

import (
	"encoding/json"
	"log"
	"net/http"
)

// Error codes used in the "error" field of error responses.
const (
	codeValidation     = "validation"      // 400
	codeNotFound       = "not_found"       // 404
	codeConflict       = "conflict"        // 409
	codeUnprocessable  = "unprocessable"   // 422
	codeNotImplemented = "not_implemented" // 501, only while routes are stubs
)

type errorBody struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Printf("write response: %v", err)
	}
}

func writeError(w http.ResponseWriter, status int, code, message string) {
	writeJSON(w, status, errorBody{Error: code, Message: message})
}
