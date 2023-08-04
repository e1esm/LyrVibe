package models

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type JWTCustomClaims struct {
	UserID   uuid.UUID `json:"id"`
	UserRole Role      `json:"role"`
	jwt.StandardClaims
}

type CachedTokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
