package service

import (
	"errors"
	"fmt"
	"github.com/e1esm/LyrVibe/auth-service/internal/models"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"math/rand"
	"time"
)

var (
	parseError       = errors.New("couldn't have extracted token")
	methodError      = errors.New("invalid signing method")
	wrongTypeError   = errors.New("token claims are of the wrong type")
	uuidParsingError = errors.New("error while parsing UUID")
)

type TokenPayload struct {
	ID       uuid.UUID   `json:"id"`
	Username string      `json:"username"`
	Role     models.Role `json:"role"`
}

type TokenManager interface {
	NewJWT(user *models.User) (string, error)
	ParseToken(string) (TokenPayload, error)
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
		Username: user.Username,
		UserRole: user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ts.accessTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})
	return token.SignedString([]byte(ts.signingKey))
}

func (ts *TokenService) ParseToken(accessToken string) (TokenPayload, error) {
	receivedToken, err := jwt.ParseWithClaims(accessToken, &models.JWTCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, methodError
		}
		return []byte(ts.signingKey), nil
	})
	if err != nil {
		return TokenPayload{}, parseError
	}
	claims, ok := receivedToken.Claims.(*models.JWTCustomClaims)
	if !ok {
		return TokenPayload{}, wrongTypeError
	}
	ID, err := uuid.Parse(claims.Id)
	if err != nil {
		return TokenPayload{}, uuidParsingError
	}
	return TokenPayload{Username: claims.Username, Role: claims.UserRole, ID: ID}, nil
}

func (ts *TokenService) NewRefreshToken() (string, error) {
	bytes := make([]byte, 32)
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	_, err := r.Read(bytes)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", bytes), nil
}
