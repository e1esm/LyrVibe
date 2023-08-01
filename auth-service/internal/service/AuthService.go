package service

import "github.com/e1esm/LyrVibe/auth-service/internal/repository"

type Service interface {
}

type AuthService struct {
	Repositories repository.Repositories
}

func NewAuthService(repositories repository.Repositories) Service {
	return &AuthService{repositories}
}
