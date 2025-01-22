package service

import (
	"errors"
	"github.com/go_web/internal/models"
	"github.com/go_web/internal/repository"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (r *AuthService) CreateUser(client models.User) error {
	return r.repo.CreateUser(client)
}

func (r *AuthService) Login(username string, password string) (models.User, error) {
	client, err := r.repo.GetUser(username, password)
	if client.Username == "" || client.Password == "" {
		err = errors.New("invalid username or password")
	}
	return client, err
}
