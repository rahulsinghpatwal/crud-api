package utils

import (
	"crud/internal/entity"
	"encoding/json"
	"errors"
	"net/http"
	"regexp"

	"github.com/dgrijalva/jwt-go"
)

type ErrorResponse struct {
	ErrorMessage string
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func JsonResponse(w http.ResponseWriter, status int, resp interface{}) {
	response, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}

var regexpEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func ValidateUser(user entity.User) error {
	if user.Firstname == "" {
		return errors.New("please provide first name")
	}
	if user.Lastname == "" {
		return errors.New("please provide last name")
	}
	if len(user.Firstname)+len(user.Lastname) > 30 {
		return errors.New("name cannot be more than 30 characters")
	}

	if len(user.Password) < 8 || len(user.Password) > 20 {
		return errors.New("password cannot be less than 8 and more than 20 ")
	}
	if len(user.Email) > 20 {
		return errors.New("email cannot be more than 20 charactes")
	}
	if !regexpEmail.Match([]byte(user.Email)) {
		return errors.New("not a valid email adderess")
	}
	return nil
}
