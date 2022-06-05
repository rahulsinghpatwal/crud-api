package repo

import (
	"crud/internal/config"
	"database/sql"
	"fmt"
)

func MigrateDb(dbname string, config *config.Config) error {
	var dbURI = fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", config.Host, config.Dbuser, config.Dbname, config.Password, config.Port)

	db, err := sql.Open(dbname, dbURI)
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users(
		id SERIAL PRIMARY KEY,
		firstName VARCHAR NOT NULL,
		lastName VARCHAR NOT NULL,
		email VARCHAR NOT NULL UNIQUE,
		password VARCHAR NOT NULL,
		dob timestamp NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  		updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  		last_accessed TIMESTAMP NOT NULL DEFAULT NOW(),
		archived BOOLEAN NOT NULL DEFAULT FALSE  
	);`)
	if err != nil {
		return err
	}

	return nil
}
