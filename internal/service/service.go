package service

import (
	"github.com/go_web/internal/models"
	"github.com/go_web/internal/repository"
)

type Authorization interface {
	CreateUser(c models.Client) error
	Login(username string, password string) (models.Client, error)
}

type Service struct {
	Authorization
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
	}
}
