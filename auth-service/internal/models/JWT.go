package models

import (
	"encoding/json"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"time"
)

type JWTCustomClaims struct {
	UserID   uuid.UUID `json:"id"`
	Username string    `json:"username"`
	UserRole Role      `json:"role"`
	jwt.StandardClaims
}

type CachedTokens struct {
	AccessToken  string        `json:"access_token"`
	AccessTTL    time.Duration `json:"access_ttl"`
	RefreshToken string        `json:"refresh_token"`
	RefreshTTL   time.Duration `json:"refresh_ttl"`
}

func (ct CachedTokens) MarshalBinary() (data []byte, err error) {
	data, err = json.Marshal(&ct)
	return
}
