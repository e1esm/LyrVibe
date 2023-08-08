package service

import (
	"github.com/e1esm/LyrVibe/auth-service/api/v1/proto"
	"github.com/e1esm/LyrVibe/gateway/internal/registrator"
	"github.com/e1esm/LyrVibe/gateway/pkg/config"
)

type AuthenticationProvider interface {
}

type AuthenticationService struct {
	client proto.AuthServiceClient
}

func NewAuthenticationService(cfg config.Config) AuthenticationProvider {
	return &AuthenticationService{client: registrator.RegisterAuthService(cfg)}
}
