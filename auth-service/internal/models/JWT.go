package models

import (
	"encoding/json"
	"github.com/e1esm/LyrVibe/auth-service/pkg/logger"
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

func (ct *CachedTokens) UnmarshalBinary(data []byte) error {
	tokens := &CachedTokens{}
	err := json.Unmarshal(data, tokens)
	if err != nil {
		logger.GetLogger().Error(err.Error())
		return err
	}
	ct.AccessTTL = tokens.AccessTTL
	ct.RefreshToken = tokens.RefreshToken
	ct.AccessToken = tokens.AccessToken
	ct.RefreshTTL = tokens.RefreshTTL
	return nil
}
