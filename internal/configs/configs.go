package configs

import "os"

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
			User:    "sumlim",
			Host:    "localhost",
			Port:    "5432",
			Pass:    "$sum-lim8876",
			Name:    "db_yalkin",
			Sslmode: "disable",
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
