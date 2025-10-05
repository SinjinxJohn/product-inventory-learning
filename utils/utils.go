package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func ParseJson(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("body is empty")
	}

	return json.NewDecoder(r.Body).Decode(payload)
}

func WriteJson(w http.ResponseWriter, status int, v any) {
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJson(w, status, map[string]string{"error": err.Error()})
}

func ValidateTokenFormat(authHeader string) (string, error) {
	parts := strings.Split(authHeader, " ")
	// fmt.Println(parts)
	if len(parts) != 2 || strings.TrimSpace(parts[0]) != "Bearer" {
		return "", fmt.Errorf("not a valid authorization format")
	}
	return parts[1], nil

}
