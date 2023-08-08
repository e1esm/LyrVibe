package service

import (
	"context"
	"github.com/e1esm/LyrVibe/auth-service/api/v1/proto"

	"github.com/e1esm/LyrVibe/gateway/internal/registrator"
	"github.com/e1esm/LyrVibe/gateway/pkg/config"
)

type AuthenticationProvider interface {
	SignUp(*proto.SignUpRequest) (*proto.SignUpResponse, error)
	Login(*proto.SignInRequest) (*proto.SignInResponse, error)
	Logout(*proto.LogoutRequest) error
}

type AuthenticationService struct {
	client proto.AuthServiceClient
}

func NewAuthenticationService(cfg config.Config) AuthenticationProvider {
	return &AuthenticationService{client: registrator.RegisterAuthService(cfg)}
}

func (as *AuthenticationService) SignUp(request *proto.SignUpRequest) (*proto.SignUpResponse, error) {
	return as.client.SignUp(context.Background(), request)
}

func (as *AuthenticationService) Login(request *proto.SignInRequest) (*proto.SignInResponse, error) {
	return as.client.SignIn(context.Background(), request)
}

func (as *AuthenticationService) Logout(request *proto.LogoutRequest) error {
	_, err := as.client.Logout(context.Background(), request)
	return err
}
