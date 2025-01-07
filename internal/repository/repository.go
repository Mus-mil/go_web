package repository

import (
	"database/sql"
	"github.com/go_web/internal/models"
)

type Authorization interface {
	CreateUser(client models.Client) error
	GetUser(username string, password string) (models.Client, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
