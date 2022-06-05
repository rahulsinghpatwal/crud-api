package repo

import (
	"crud/internal/entity"
	"crud/utils"
	"errors"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func (r Repository) CreateUser(user entity.User) (*utils.UserResponse, error) {
	// check email is taken or not
	var count = 0
	var response utils.UserResponse
	err := r.db.QueryRow(`Select email from users where email=$1`, user.Email).Scan(&count)
	if err != nil {
		if count == 1 {
			return nil, errors.New("invalid email address")
		}
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		zap.S().Error(err)
		return nil, err
	}

	user.Password = string(bytes)

	_, err = r.db.Exec(`Insert into users(firstname, lastname, email, password, dob) values(
							$1, $2, $3, $4, $5)`, user.Firstname, user.Lastname, user.Email, user.Password,
		user.Dob)

	if err != nil {
		zap.S().Error(err)
		return nil, errors.New("internal server error")
	}

	mapResponse(user, &response)
	var user_id string
	err = r.db.QueryRow(`select id from users where email = $1`, user.Email).Scan(&user_id)
	if err != nil {
		zap.S().Error(err)
		return nil, errors.New("internal server error")
	}
	response.Id = user_id

	return &response, nil
}

func mapResponse(user entity.User, response *utils.UserResponse) {
	response.Firstname = user.Firstname
	response.Lastname = user.Lastname
	response.Email = user.Email
	response.Dob = user.Dob
}
