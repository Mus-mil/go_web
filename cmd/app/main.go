package main

import (
	"database/sql"
	"github.com/go_web/internal/configs"
	"github.com/go_web/internal/handlers"
	"github.com/go_web/internal/repository"
	"github.com/go_web/internal/service"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := godotenv.Load("/home/sum-lim/go/src/github.com/go_web/.env"); err != nil {
		log.Print("No .env file found")
	}

	if err := initConfigs(); err != nil {
		log.Print("No .env file found")
	}

	cfg := configs.NewConfigs()

	db := repository.InitDB(cfg.Postgres)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal("Not closer DB")
		}
	}(db)

	repo := repository.NewRepository(db)
	serv := service.NewService(repo)
	handler := handlers.NewHandler(serv)

	srv := new(handlers.Server)
	if err := srv.Run(handlers.RegisterRoutes(handler)); err != nil {
		log.Fatal("server not running")
	}
	log.Print("server run in http://localhost:8080")
}

func initConfigs() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
