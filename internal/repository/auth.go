package repository

import (
	"database/sql"
	"github.com/go_web/internal/domain"
)

type AuthPostgres struct {
	db *sql.DB
}

func NewAuthPostgres(db *sql.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (repo *AuthPostgres) CreateUser(client domain.Client) error {
	_, err := repo.db.Exec("INSERT INTO users (name, username, password) VALUES ($1, $2, $3)",
		client.Name, client.Username, client.Password)
	if err != nil {
		return err
	}
	return nil
}

func (repo *AuthPostgres) GetUser(username string, password string) (domain.Client, error) {
	var client domain.Client
	row := repo.db.QueryRow("SELECT id, name, username, password FROM users WHERE username = $1 AND password = $2", username, password)
	err := row.Scan(&client.ID, &client.Name, &client.Username, &client.Password)
	if err != nil {
		return domain.Client{}, err
	}
	return client, nil
}
