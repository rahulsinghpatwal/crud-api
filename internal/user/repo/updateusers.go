package repo

import (
	"crud/utils"
	"errors"
	"time"

	"go.uber.org/zap"
)

func (r Repository) UpdateUser(userID string, user utils.UpdateUserRequest) (*utils.UpdateResponse, error) {

	var resp utils.UpdateResponse
	var prevDetails utils.UpdateResponse

	err := r.db.QueryRow(`select firstname, lastname, dob from users where id = $1 and archived = false`, userID).Scan(&prevDetails.Firstname, &prevDetails.Lastname, &prevDetails.Dob)

	zap.S().Info(prevDetails)
	if err != nil {
		zap.S().Error(err)
		return nil, errors.New("internal server error")
	}
	mapping(&user, prevDetails)
	err = r.db.QueryRow(`Update users set firstname = $1, lastname=$2, dob=$3 , updated_at =$4 where id=$5 and archived = false returning firstname, lastname, dob`, user.Firstname, user.Lastname, user.Dob, time.Now(), userID).Scan(&resp.Firstname, &resp.Lastname, &resp.Dob)

	zap.S().Info("user:", user)
	if err != nil {
		zap.S().Error(err)
		return nil, errors.New("internal server error")
	}

	if err != nil {
		zap.S().Error(err)
		return nil, errors.New("internal server error")
	}

	return &resp, nil

}

func mapping(user *utils.UpdateUserRequest, prevDetails utils.UpdateResponse) {
	if user.Firstname == "" {
		user.Firstname = prevDetails.Firstname
	}
	if user.Lastname == "" {
		user.Lastname = prevDetails.Lastname
	}

	if user.Dob.IsZero() {
		user.Dob = prevDetails.Dob
	}
}
