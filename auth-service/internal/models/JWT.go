package models

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"time"
)

type JWTCustomClaims struct {
	UserID   uuid.UUID `json:"id"`
	UserRole Role      `json:"role"`
	jwt.StandardClaims
}

type CachedTokens struct {
	AccessToken  string        `json:"access_token"`
	AccessTTL    time.Duration `json:"access_ttl"`
	RefreshToken string        `json:"refresh_token"`
	RefreshTTL   time.Duration `json:"refresh_ttl"`
}
