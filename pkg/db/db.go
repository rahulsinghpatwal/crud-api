package db

import (
	"crud/internal/config"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var err error
var db *sql.DB

func CreateConnection(config *config.Config) (*sql.DB, error) {
	var dbURI = fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", config.Host, config.Dbuser, config.Dbname, config.Password, config.Port)

	db, err = sql.Open("postgres", dbURI)
	if err != nil {
		return nil, err
	}
	err = db.Ping()

	if err != nil {
		return nil, err
	}
	return db, nil
}
