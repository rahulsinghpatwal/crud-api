package middleware

import (
	"crud/internal/config"
	"crud/utils"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func Authorize(next func(w http.ResponseWriter, req *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		jwtKey := config.LoadJwt()
		key := jwtKey.Key
		authHeader := strings.Split(req.Header.Get("Authorization"), " ")
		if len(authHeader) != 2 {
			http.Error(w, "not authorized", http.StatusUnauthorized)
			return
		}

		if authHeader[1] == "" {
			http.Error(w, "not authorized", http.StatusUnauthorized)
			return
		} else {
			jwtToken := authHeader[1]
			claims := &utils.Claims{}

			tkn, err := jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(key), nil
			})
			if err != nil {
				fmt.Println(err)
				http.Error(w, "not authorized", http.StatusUnauthorized)
				return
			}

			if tkn.Valid {
				next(w, req)
			}
		}

	})
}
