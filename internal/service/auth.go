package service

import (
	"crypto/sha1"
	"encoding/hex"
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
	client.Password = r.generatePasswordHash(client.Password)
	return r.repo.CreateUser(client)
}

func (r *AuthService) LoginUser(username string, password string) (models.User, error) {
	client, err := r.repo.GetUser(username, password)
	if client.Username == "" || client.Password == "" {
		err = errors.New("invalid username or password")
	}
	return client, err
}

func (r *AuthService) generatePasswordHash(password string) string {
	passwordHash := sha1.New()
	passwordHash.Write([]byte(password))

	return hex.EncodeToString(passwordHash.Sum(nil))
}
