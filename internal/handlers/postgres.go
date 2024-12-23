package handlers

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
)

type postgresConfig struct {
	user    string
	host    string
	port    string
	pass    string
	name    string
	sslmode string
}

func (cfg postgresConfig) connectDB() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.host, cfg.port, cfg.user, cfg.pass, cfg.name, cfg.sslmode)
}

func editUsersTable(db *sql.DB, client Client) error {
	_, err := db.Exec("INSERT INTO users (name, username, password) VALUES ($1, $2, $3)",
		client.Name, client.Username, client.Password)
	if err != nil {
		return err
	}
	return nil
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	log.Println("")
	return viper.ReadInConfig()
}

func initDB() *sql.DB {
	if err := initConfig(); err != nil {
		log.Fatal("error initializing configs")
	}
	cfg := postgresConfig{
		user:    viper.GetString("db.username"),
		host:    viper.GetString("db.host"),
		port:    viper.GetString("db.port"),
		pass:    viper.GetString("db.password"),
		name:    viper.GetString("db.db_name"),
		sslmode: viper.GetString("db.sslmode"),
	}
	db, err := sql.Open("postgres", cfg.connectDB())
	if err != nil {
		log.Println(cfg.connectDB())
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Ошибка при проверке соединения: %v", err)
	}
	return db
}

func closeDB(db *sql.DB) {
	if err := db.Close(); err != nil {
		log.Fatal(err)
	}
}

func getUserFromUsername(db *sql.DB, username string) (Client, error) {
	var client Client
	row := db.QueryRow("SELECT id, name, username, password FROM users WHERE username = $1", username)
	err := row.Scan(&client.ID, &client.Name, &client.Username, &client.Password)
	return client, err
}
