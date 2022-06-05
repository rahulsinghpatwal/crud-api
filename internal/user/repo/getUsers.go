package repo

import (
	"crud/utils"
	"errors"
	"time"

	"go.uber.org/zap"
)

func (r Repository) ListUsers(l int, offset int) (*[]utils.UserResponse, error) {
	var users []utils.UserResponse
	var user utils.UserResponse

	rows, err := r.db.Query(`select id, firstname, lastname, email, dob from users where archived=false Order by id offset $1 limit $2`, offset, l)

	if err != nil {

		zap.S().Error(err)
		return nil, errors.New("internal server error")
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Firstname, &user.Lastname, &user.Email, &user.Dob)
		if err != nil {
			zap.S().Error(err)
			return nil, errors.New("internal server error")
		}

		_, err = r.db.Exec(`update users set last_accessed = $1 where email=$2`, time.Now(), user.Email)
		if err != nil {
			zap.S().Error(err)
			return nil, errors.New("internal server error")
		}
		users = append(users, user)
	}
	return &users, nil

}
