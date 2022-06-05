package main

import (
	"crud/internal/config"
	"crud/internal/user"
	"crud/internal/user/goroutine"
	"crud/internal/user/repo"
	"crud/logger"
	"crud/pkg/db"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"go.uber.org/zap"
)

func main() {
	logger.InitLogger()
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("err loading env file")
	}

	config := config.Load()
	tickerTime := os.Getenv("TICKERTIME")
	db, err := db.CreateConnection(config)
	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
	})

	err = repo.MigrateDb("postgres", config)
	if err != nil {
		zap.S().Error(err)
		return
	}
	repo := repo.NewRepository(db)
	service := user.NewService(repo)
	user.RegisterHandler(config, router, service)

	handler := c.Handler(router)
	t, err := strconv.Atoi(tickerTime)
	if err != nil {
		zap.S().Error(err)
	}

	duration := time.Duration(t) * time.Minute
	tk := time.NewTicker(duration)
	go goroutine.Check(tk)

	err = http.ListenAndServe(":8080", handler)
	if err != nil {
		zap.S().Error(err)
	}
}
