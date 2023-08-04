package service

import (
	"context"
	"database/sql"
	"github.com/e1esm/LyrVibe/auth-service/internal/models"
	"github.com/e1esm/LyrVibe/auth-service/internal/repository"
	"github.com/e1esm/LyrVibe/auth-service/pkg/logger"
	"github.com/joho/godotenv"
	"os"
	"time"
)

const (
	defaultTTL = 100
)

type Service interface {
	SaveUser(context.Context, *models.User) error
	GetUser(context.Context, string, string) (*models.User, error)
	CreateSession(context.Context, *models.User) (models.Tokens, error)
}

type AuthService struct {
	Repositories repository.Repositories
	TokenService TokenManager
	ttl          time.Duration
}

func init() {
	err := godotenv.Load("jwt.env")
	if err != nil {
		logger.Logger.Error("Couldn't have loaded file with jwt env variables")
	}
}

func NewAuthService(repositories repository.Repositories, serviceBuilder TokenServiceBuilder) Service {
	ttl, err := time.ParseDuration(os.Getenv("TTL"))
	if err != nil {
		ttl = defaultTTL
	}
	manager := serviceBuilder.WithSigningKey(os.Getenv("SIGNING_KEY")).WithTTL(ttl).Build()
	return &AuthService{repositories, manager, ttl}
}

func (as *AuthService) SaveUser(ctx context.Context, user *models.User) error {
	err := as.Repositories.MainRepository.Add(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (as *AuthService) GetUser(ctx context.Context, username, password string) (*models.User, error) {
	user := as.Repositories.MainRepository.GetOne(ctx, username, password)
	if user == nil {
		return nil, sql.ErrNoRows
	}
	return user, nil
}

func (as *AuthService) CreateSession(ctx context.Context, user *models.User) (models.Tokens, error) {
	return models.Tokens{}, nil
}
