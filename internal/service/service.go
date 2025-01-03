package service

import (
	"github.com/go_web/internal/domain"
	"github.com/go_web/internal/repository"
)

type Authorization interface {
	CreateUser(c domain.Client) error
	Login(username string, password string) (domain.Client, error)
}

type Service struct {
	Authorization
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
	}
}
