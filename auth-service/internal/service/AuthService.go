package service

import (
	"context"
	"database/sql"
	"github.com/e1esm/LyrVibe/auth-service/internal/models"
	"github.com/e1esm/LyrVibe/auth-service/internal/repository"
	"github.com/e1esm/LyrVibe/auth-service/pkg/logger"
	"github.com/google/uuid"
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
	CreateSession(context.Context, *models.User) (models.CachedTokens, error)
	GetSessionCredentials(context.Context, uuid.UUID) (models.CachedTokens, error)
}

type AuthService struct {
	Repositories repository.Repositories
	TokenService TokenManager
	accessTTL    time.Duration
	refreshTTL   time.Duration
}

func init() {
	err := godotenv.Load("jwt.env")
	if err != nil {
		logger.Logger.Error("Couldn't have loaded file with jwt env variables")
	}
}

func NewAuthService(repositories repository.Repositories, serviceBuilder TokenServiceBuilder) Service {
	accessTTL, err := time.ParseDuration(os.Getenv("ACCESS_TTL"))
	refreshTTL, err := time.ParseDuration(os.Getenv("REFRESH_TTL"))
	if err != nil {
		accessTTL = defaultTTL
	}
	manager := serviceBuilder.WithSigningKey(os.Getenv("SIGNING_KEY")).WithTTL(accessTTL).Build()
	return &AuthService{repositories, manager, accessTTL, refreshTTL}
}

func (as *AuthService) GetSessionCredentials(ctx context.Context, id uuid.UUID) (models.CachedTokens, error) {
	tokens, err := as.Repositories.SessionRepository.Get(ctx, id)
	if err != nil {
		return models.CachedTokens{}, err
	}
	return tokens, nil
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

func (as *AuthService) CreateSession(ctx context.Context, user *models.User) (models.CachedTokens, error) {
	jwtToken, err := as.TokenService.NewJWT(user)
	if err != nil {
		return models.CachedTokens{}, err
	}
	refreshToken, err := as.TokenService.NewRefreshToken()
	if err != nil {
		return models.CachedTokens{}, err
	}

	tokens := models.CachedTokens{
		AccessTTL:    as.accessTTL,
		RefreshTTL:   as.refreshTTL,
		AccessToken:  jwtToken,
		RefreshToken: refreshToken,
	}

	return as.Repositories.SessionRepository.Add(ctx, user, tokens)
}
