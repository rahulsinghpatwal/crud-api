package config

import (
	"os"
)

type Config struct {
	DBdriver string
	Host     string
	Port     string
	Dbuser   string
	Dbname   string
	Password string
}

type JwtKey struct {
	Key string
}

func Load() *Config {
	config := &Config{}
	config.Host = os.Getenv("HOST")
	config.Dbname = os.Getenv("DBNAME")
	config.Dbuser = os.Getenv("DBUSER")
	config.Port = os.Getenv("PORT")
	config.Password = os.Getenv("PASSWORD")

	return config
}

func LoadJwt() JwtKey {
	Jwt := JwtKey{}
	Jwt.Key = os.Getenv("JwtKey")
	return Jwt
}
