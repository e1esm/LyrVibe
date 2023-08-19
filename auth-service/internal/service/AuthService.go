package service

import (
	"context"
	"database/sql"
	"errors"
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

var (
	wasAlreadyCached = errors.New("already logged in")
)

type Service interface {
	Logout(context.Context, string) error
	UpdateRole(context.Context, uuid.UUID, models.Role) error
	SessionManager
	UserManager
}

type UserManager interface {
	SaveUser(context.Context, *models.User) error
	GetUser(context.Context, string, string) (*models.User, error)
	GetCredentials(string) (TokenPayload, error)
}

type SessionManager interface {
	CreateSession(context.Context, *models.User) (models.CachedTokens, error)
	GetSessionCredentials(context.Context, string) (models.CachedTokens, error)
	UpdateSession(context.Context, models.CachedTokens) (bool, error)
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
		logger.Logger.Info(err.Error())
		accessTTL = defaultTTL
	}
	manager := serviceBuilder.WithSigningKey(os.Getenv("SIGNING_KEY")).WithTTL(accessTTL).Build()
	return &AuthService{repositories, manager, accessTTL, refreshTTL}
}

func (as *AuthService) GetSessionCredentials(ctx context.Context, refreshToken string) (models.CachedTokens, error) {
	tokens, err := as.Repositories.SessionRepository.Get(ctx, refreshToken)
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

	wasAdded, err := as.Repositories.SessionRepository.Add(ctx, tokens)
	if !wasAdded {
		return models.CachedTokens{}, wasAlreadyCached
	}
	return tokens, nil
}

func (as *AuthService) Logout(ctx context.Context, accessToken string) error {
	payload, err := as.TokenService.ParseToken(accessToken)
	if err != nil {
		return err
	}
	err = as.Repositories.SessionRepository.Delete(ctx, payload.ID)
	return err
}

func (as *AuthService) UpdateRole(ctx context.Context, id uuid.UUID, role models.Role) error {
	return as.Repositories.MainRepository.UpdateRole(ctx, id, role)
}

func (as *AuthService) GetCredentials(accessToken string) (TokenPayload, error) {
	return as.TokenService.ParseToken(accessToken)
}

func (as *AuthService) UpdateSession(ctx context.Context, tokens models.CachedTokens) (bool, error) {
	return as.Repositories.SessionRepository.Add(ctx, tokens)
}
