package service

import (
	"errors"
	"fmt"
	"github.com/e1esm/LyrVibe/auth-service/internal/models"

	"github.com/e1esm/LyrVibe/auth-service/pkg/logger"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"math/rand"
	"time"
)

var (
	parseError     = errors.New("couldn't have extracted token")
	methodError    = errors.New("invalid signing method")
	wrongTypeError = errors.New("token claims are of the wrong type")
)

type TokenManager interface {
	NewJWT(user *models.User) (string, error)
	ParseToken(string) (uuid.UUID, models.Role, error)
	NewRefreshToken() (string, error)
}

type TokenService struct {
	signingKey string
	accessTTL  time.Duration
}

type TokenServiceBuilder struct {
	TokenService TokenService
}

func NewTokenServiceBuilder() *TokenServiceBuilder {
	return &TokenServiceBuilder{TokenService: TokenService{}}
}

func (tsb *TokenServiceBuilder) WithSigningKey(signingKey string) *TokenServiceBuilder {
	tsb.TokenService.signingKey = signingKey
	return tsb
}

func (tsb *TokenServiceBuilder) WithTTL(TTL time.Duration) *TokenServiceBuilder {
	tsb.TokenService.accessTTL = TTL
	return tsb
}

func (tsb *TokenServiceBuilder) Build() TokenManager {
	return &tsb.TokenService
}

func (ts *TokenService) NewJWT(user *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, models.JWTCustomClaims{
		UserID:   user.ID,
		UserRole: user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ts.accessTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})
	return token.SignedString([]byte(ts.signingKey))
}

func (ts *TokenService) ParseToken(accessToken string) (uuid.UUID, models.Role, error) {
	receivedToken, err := jwt.ParseWithClaims(accessToken, &models.JWTCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, methodError
		}
		return []byte(ts.signingKey), nil
	})
	if err != nil {
		return uuid.UUID{}, "", parseError
	}
	claims, ok := receivedToken.Claims.(*models.JWTCustomClaims)
	if !ok {
		return uuid.UUID{}, "", wrongTypeError
	}
	return claims.UserID, claims.UserRole, nil
}

func (ts *TokenService) NewRefreshToken() (string, error) {
	bytes := make([]byte, 32)
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	writtenLen, err := r.Read(bytes)
	logger.Logger.Info("Written length", zap.Int("len", writtenLen))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", bytes), nil
}
