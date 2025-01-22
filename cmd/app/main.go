package main

import (
	"github.com/go_web/internal/app"
	"github.com/go_web/internal/configs"
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

	app.RunServer(cfg)
}

func initConfigs() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
