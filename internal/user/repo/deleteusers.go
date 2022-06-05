package repo

import (
	"crud/utils"
	"errors"
	"time"

	"go.uber.org/zap"
)

func (r Repository) DeleteUser(userID string) (*utils.DeleteResponse, error) {

	var response utils.DeleteResponse
	var count = 0

	err := r.db.QueryRow(`Select count(id) from users where id = $1 and archived = false`, userID).Scan(&count)
	if err != nil {
		zap.S().Error(err)
		return nil, errors.New("internal server error")
	}

	if count == 0 {
		return nil, errors.New("user does not exists in the system")
	}

	_, err = r.db.Exec(`Update users set archived = true, last_accessed=$1 where id = $2`, time.Now(), userID)
	if err != nil {
		zap.S().Error(err)
		return nil, errors.New("internal server error")
	}

	response.Message = "User deleted successfully"
	return &response, nil

}
