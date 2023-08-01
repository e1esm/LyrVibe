package service

import (
	"context"
	"github.com/e1esm/LyrVibe/auth-service/internal/models"
	"github.com/e1esm/LyrVibe/auth-service/internal/repository"
)

type Service interface {
	SaveUser(context.Context, *models.User) error
}

type AuthService struct {
	Repositories repository.Repositories
}

func NewAuthService(repositories repository.Repositories) Service {
	return &AuthService{repositories}
}

func (as *AuthService) SaveUser(ctx context.Context, user *models.User) error {
	err := as.Repositories.MainRepository.Add(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
