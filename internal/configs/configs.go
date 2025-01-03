package configs

import (
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	Postgres PostgresConfig
}

type PostgresConfig struct {
	User    string
	Host    string
	Port    string
	Pass    string
	Name    string
	Sslmode string
}

func NewConfigs() Config {
	cfg := Config{
		Postgres: PostgresConfig{
			User:    viper.GetString("db.user"),
			Host:    viper.GetString("db.host"),
			Port:    viper.GetString("db.port"),
			Pass:    getEnv("DB_PASSWORD", ""),
			Name:    viper.GetString("db.name"),
			Sslmode: viper.GetString("db.sslmode"),
		},
	}
	return cfg
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
