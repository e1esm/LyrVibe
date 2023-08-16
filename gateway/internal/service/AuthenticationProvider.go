package service

import (
	"context"
	"github.com/e1esm/LyrVibe/auth-service/api/v1/proto"
	"github.com/e1esm/LyrVibe/gateway/internal/registrator"
	"github.com/e1esm/LyrVibe/gateway/pkg/config"
	"github.com/e1esm/LyrVibe/gateway/pkg/logger"
)

type AuthenticationProvider interface {
	SignUp(*proto.SignUpRequest) (*proto.SignUpResponse, error)
	Login(*proto.SignInRequest) (*proto.SignInResponse, error)
	Logout(*proto.LogoutRequest) error
	Verify(request *proto.VerificationRequest) (*proto.VerificationResponse, error)
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

func (as *AuthenticationService) Verify(request *proto.VerificationRequest) (*proto.VerificationResponse, error) {
	resp, err := as.client.Verification(context.Background(), request)
	if err != nil {
		logger.Logger.Error(err.Error())
		return nil, err
	}
	return resp, nil
}
