// utils/utils.go
package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

// RespondJSON sends a JSON response with the given status code and data
func RespondJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Printf("Failed to encode JSON response: %v", err)
	}
}

// ParseJSONRequest parses the JSON request body into the provided data structure
func ParseJSONRequest(r *http.Request, data interface{}) error {
	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		log.Printf("Failed to parse JSON request: %v", err)
	}
	return err
}
