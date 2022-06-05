package repo

import (
	"database/sql"
	"errors"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return Repository{db}
}

func (r Repository) GetUserByEmail(email string, password string) error {
	var dbPass string
	err := r.db.QueryRow(`select password from users where email = $1 and archived = false`, email).Scan(&dbPass)
	if err != nil {
		zap.S().Error(err)
		return errors.New("internal server error")
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbPass), []byte(password))
	if err != nil {
		zap.S().Error(err)
		return errors.New("internal server error")
	}

	return nil

}
