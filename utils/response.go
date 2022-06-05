package utils

import "time"

type UserResponse struct {
	Id        string    `json:"id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Email     string    `json:"email"`
	Dob       time.Time `json:"dob"`
}

type DeleteResponse struct {
	Message string `json:"message"`
}

type UpdateResponse struct {
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Dob       time.Time `json:"dob"`
}

type JwtTokenResponse struct {
	Token string
}

type SoftDeletedRecord struct {
	Id        string `json:"id"`
	Firstname string `json:"firstname"`
}
