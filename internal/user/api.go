package user

import (
	"crud/internal/config"
	"crud/internal/entity"
	"crud/internal/model"
	"crud/internal/user/middleware"
	"crud/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type resource struct {
	config  *config.Config
	service model.Service
}

func RegisterHandler(config *config.Config, r *mux.Router, service model.Service) {
	res := resource{config, service}
	r.HandleFunc("/login", res.Login).Methods("POST")
	r.HandleFunc("/user/{id}", res.GetUserById).Methods("GET")
	r.HandleFunc("/list", res.ListUsers).Methods("GET")
	r.HandleFunc("/user", res.CreateUser).Methods("POST")
	r.Handle("/user/{id}", middleware.Authorize(res.DeleteUser)).Methods("DELETE")
	r.Handle("/user/{id}", middleware.Authorize(res.UpdateUser)).Methods("PUT")
}

func (res resource) Login(w http.ResponseWriter, req *http.Request) {

	var cred utils.LoginCredentials
	var resp utils.JwtTokenResponse
	err := utils.DecodeRequest(req, &cred)
	if err != nil {
		zap.S().Error(err)
		utils.JsonResponse(w, http.StatusBadRequest, utils.ErrorResponse{
			ErrorMessage: err.Error(),
		})
		return
	}

	if cred.Email == "" || cred.Password == "" {
		utils.JsonResponse(w, http.StatusBadRequest, utils.ErrorResponse{
			ErrorMessage: "Bad request",
		})
		return
	}

	err = res.service.GetUserByEmail(cred.Email, cred.Password)
	if err != nil {
		zap.S().Error(err)
		utils.JsonResponse(w, http.StatusBadRequest, utils.ErrorResponse{
			ErrorMessage: "Bad Request",
		})
		return
	}
	jwtKey := config.LoadJwt()
	key := []byte(jwtKey.Key)

	expirationTime := time.Now().Add(30 * time.Minute).Unix()

	claims := &utils.Claims{
		Email: cred.Email,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(key)

	if err != nil {
		zap.S().Error(err)
		utils.JsonResponse(w, http.StatusBadRequest, utils.ErrorResponse{
			ErrorMessage: "Bad Request",
		})
		return
	}

	resp.Token = tokenString

	utils.JsonResponse(w, http.StatusOK, resp)

}

func (res resource) GetUserById(w http.ResponseWriter, req *http.Request) {
	userId := mux.Vars(req)["id"]
	if userId == "" {
		utils.JsonResponse(w, http.StatusBadRequest, utils.ErrorResponse{
			ErrorMessage: "Id is required",
		})
		return
	}

	resp, err := res.service.GetUserById(userId)

	if err != nil {
		zap.S().Error(err)
		utils.JsonResponse(w, http.StatusInternalServerError, utils.ErrorResponse{
			ErrorMessage: err.Error(),
		})
		return
	}

	utils.JsonResponse(w, http.StatusOK, resp)

}

func (res resource) ListUsers(w http.ResponseWriter, req *http.Request) {
	page := req.URL.Query().Get("page")
	limit := req.URL.Query().Get("limit")
	var p, l int
	var err error
	if page == "" {
		p = 1
	} else {
		p, err = strconv.Atoi(page)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
	}
	if limit == "" {
		l = 5
	} else {
		l, err = strconv.Atoi(limit)
		if err != nil {
			zap.S().Error(err)
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
	}

	offset := (p - 1) * l

	resp, err := res.service.ListUsers(l, offset)
	if err != nil {
		zap.S().Error(err)
		utils.JsonResponse(w, http.StatusInternalServerError, utils.ErrorResponse{
			ErrorMessage: err.Error(),
		})
		return
	}

	utils.JsonResponse(w, http.StatusOK, resp)

}

func (res resource) CreateUser(w http.ResponseWriter, req *http.Request) {
	var user entity.User

	err := utils.DecodeRequest(req, &user)
	if err != nil {
		zap.S().Error(err)
		zap.S().Error(err)
		utils.JsonResponse(w, http.StatusBadRequest, utils.ErrorResponse{
			ErrorMessage: err.Error(),
		})
		return
	}

	err = utils.ValidateUser(user)
	if err != nil {
		zap.S().Error(err)
		utils.JsonResponse(w, http.StatusBadRequest, utils.ErrorResponse{
			ErrorMessage: err.Error(),
		})
		return
	}

	resp, err := res.service.CreateUser(user)
	if err != nil {
		zap.S().Error(err)
		utils.JsonResponse(w, http.StatusInternalServerError, utils.ErrorResponse{
			ErrorMessage: err.Error(),
		})
		return
	}

	utils.JsonResponse(w, http.StatusOK, resp)
}

func (res resource) DeleteUser(w http.ResponseWriter, req *http.Request) {
	userId := mux.Vars(req)["id"]
	if userId == "" {
		utils.JsonResponse(w, http.StatusBadRequest, utils.ErrorResponse{
			ErrorMessage: "Id is required",
		})
		return
	}

	resp, err := res.service.DeleteUser(userId)
	if err != nil {
		zap.S().Error(err)
		utils.JsonResponse(w, http.StatusInternalServerError, utils.ErrorResponse{
			ErrorMessage: err.Error(),
		})
		return
	}

	utils.JsonResponse(w, http.StatusOK, resp)

}

func (res resource) UpdateUser(w http.ResponseWriter, req *http.Request) {
	userId := mux.Vars(req)["id"]
	if userId == "" {
		utils.JsonResponse(w, http.StatusBadRequest, utils.ErrorResponse{
			ErrorMessage: "Id is required",
		})
		return
	}

	var user utils.UpdateUserRequest
	err := utils.DecodeRequest(req, &user)
	if err != nil {
		zap.S().Error(err)
		utils.JsonResponse(w, http.StatusBadRequest, utils.ErrorResponse{
			ErrorMessage: err.Error(),
		})
		return
	}

	resp, err := res.service.UpdateUser(userId, user)

	if err != nil {
		zap.S().Error(err)
		utils.JsonResponse(w, http.StatusInternalServerError, utils.ErrorResponse{
			ErrorMessage: err.Error(),
		})
		return
	}

	utils.JsonResponse(w, http.StatusOK, resp)

}
