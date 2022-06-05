package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func DecodeRequest(r *http.Request, req interface{}) error {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	if err := decoder.Decode(req); err != nil {
		return fmt.Errorf("failed to decode request: %w", err)
	}
	return nil
}

type UpdateUserRequest struct {
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Dob       time.Time `json:"dob"`
}

type LoginCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
