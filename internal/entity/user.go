package entity

import "time"

type User struct {
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Dob       time.Time `json:"dob"`
}
