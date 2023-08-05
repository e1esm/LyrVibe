package server

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/e1esm/LyrVibe/auth-service/api/v1/proto"
	"github.com/e1esm/LyrVibe/auth-service/internal/models"
	"github.com/e1esm/LyrVibe/auth-service/internal/service"
	"github.com/e1esm/LyrVibe/auth-service/pkg/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type RequestStatus string

const (
	Success RequestStatus = "OK"
	Fail    RequestStatus = "FAIL"
)

const (
	SessionCreationError = "Failed to create a session"
	CreatingModelError   = "User with nickname of %s cannot be created"
	BadRequestError      = "bad request"
	SaveError            = "Error while saving %s"
	InternalError        = "internal error occurred while operating on the provided input"
	NoUserFound          = "User with username of %s wasn't found in the database or password was incorrect: %s"
)

type Server struct {
	Server      *grpc.Server
	AuthService service.Service
	proto.UnimplementedAuthServiceServer
}

func (s *Server) Identify(ctx context.Context, request *proto.IdentifyRequest) (*proto.IdentifyResponse, error) {
	return nil, nil
}

func (s *Server) SignUp(ctx context.Context, request *proto.SignUpRequest) (*proto.SignUpResponse, error) {
	logger.Logger.Info("User: ", zap.String("user", fmt.Sprintf("%v", request)))
	user := models.NewUser(request)
	if user == nil {
		return &proto.SignUpResponse{
			Username: request.Username,
			Status: &proto.RequestStatus{
				RequestStatus: string(Fail),
				ErrorMessage:  fmt.Sprintf(CreatingModelError, request.Username),
			},
		}, errors.New(BadRequestError)
	}

	err := s.AuthService.SaveUser(ctx, user)
	if err != nil {
		logger.Logger.Error(err.Error())
		return &proto.SignUpResponse{Username: request.Username,
			Status: &proto.RequestStatus{
				RequestStatus: string(Fail),
				ErrorMessage:  fmt.Sprintf(SaveError, request.Username)},
		}, errors.New(InternalError)
	}

	return &proto.SignUpResponse{
		Username: request.Username,
		Status: &proto.RequestStatus{
			RequestStatus: string(Success),
			ErrorMessage:  "",
		},
	}, nil
}

func (s *Server) SignIn(ctx context.Context, request *proto.SignInRequest) (*proto.SignInResponse, error) {
	user, err := s.AuthService.GetUser(ctx, request.Username, request.Password)
	if err == sql.ErrNoRows {
		return &proto.SignInResponse{
			Tokens: &proto.CachedTokens{
				AccessToken:  "",
				RefreshToken: "",
			},
			Status: &proto.RequestStatus{
				RequestStatus: string(Fail),
				ErrorMessage:  fmt.Sprintf(NoUserFound, request.Username, request.Password),
			},
		}, err
	}

	cachedTokens, err := s.AuthService.CreateSession(ctx, user)
	if err != nil {
		return &proto.SignInResponse{
			Tokens: &proto.CachedTokens{
				AccessToken:  "",
				RefreshToken: "",
			},
			Status: &proto.RequestStatus{
				RequestStatus: string(Fail),
				ErrorMessage:  SessionCreationError,
			},
		}, err
	}

	return &proto.SignInResponse{
		Tokens: &proto.CachedTokens{
			AccessToken:  cachedTokens.AccessToken,
			RefreshToken: cachedTokens.RefreshToken,
		},
		Status: &proto.RequestStatus{
			RequestStatus: string(Success),
			ErrorMessage:  "",
		},
	}, nil
}
