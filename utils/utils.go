package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ParseJSON(r *http.Request, payload any) error {
	defer r.Body.Close()

	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}

	return json.NewDecoder(r.Body).Decode(payload)
}

func WriteJSON(w http.ResponseWriter, status int, value any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(value)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"errors": err.Error()})
}
