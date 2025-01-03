package repository

import (
	"database/sql"
	"github.com/go_web/internal/domain"
)

type Authorization interface {
	CreateUser(client domain.Client) error
	GetUser(username string, password string) (domain.Client, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
