package repo

import (
	"crud/utils"
	"errors"
	"time"

	"go.uber.org/zap"
)

func (r Repository) GetUserById(userId string) (*utils.UserResponse, error) {
	var resp utils.UserResponse

	err := r.db.QueryRow(`select id, firstname, lastname, email, dob from users where id = $1 and archived = false`, userId).Scan(&resp.Id, &resp.Firstname, &resp.Lastname, &resp.Email, &resp.Dob)

	if err != nil {
		zap.S().Error(err)
		return nil, errors.New("user does not exists in the system")
	}

	_, err = r.db.Exec(`Update users set last_accessed =$1 where id=$2`, time.Now(), userId)
	if err != nil {
		zap.S().Error(err)
		return nil, errors.New("internal server error")
	}

	return &resp, nil
}
