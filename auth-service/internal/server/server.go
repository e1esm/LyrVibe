package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/e1esm/LyrVibe/auth-service/api/v1/proto"
	"github.com/e1esm/LyrVibe/auth-service/internal/models"
	"github.com/e1esm/LyrVibe/auth-service/internal/service"
	"github.com/e1esm/LyrVibe/auth-service/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type RequestStatus string

const (
	Success RequestStatus = "OK"
	Fail    RequestStatus = "FAIL"
)

const (
	CreatingModelError = "User with nickname of %s cannot be created"
	BadRequestError    = "bad request"
	SaveError          = "Error while saving %s"
	InternalError      = "internal error occurred while operating on the provided input"
)

type Server struct {
	Server      *grpc.Server
	AuthService service.Service
	proto.UnimplementedAuthServiceServer
}

func (s *Server) SignUp(ctx context.Context, request *proto.SignUpRequest) (*proto.SignUpResponse, error) {
	user := models.NewUser(request)

	if user == nil {
		return &proto.SignUpResponse{
			Username:     request.Username,
			SignupStatus: string(Fail),
			ErrorMessage: fmt.Sprintf(CreatingModelError, request.Username),
		}, errors.New(BadRequestError)
	}

	err := s.AuthService.SaveUser(ctx, user)
	if err != nil {
		logger.Logger.Error(err.Error())
		return &proto.SignUpResponse{Username: request.Username,
			SignupStatus: string(Fail),
			ErrorMessage: fmt.Sprintf(SaveError, request.Username)}, errors.New(InternalError)
	}

	return &proto.SignUpResponse{
		Username:     request.Username,
		SignupStatus: string(Success),
		ErrorMessage: "",
	}, nil
}

func (s *Server) SignIn(ctx context.Context, request *proto.SignInRequest) (*proto.SignInResponse, error) {
	return nil, nil
}

func (s *Server) RefreshToken(ctx context.Context, request *proto.RefreshRequest) (*emptypb.Empty, error) {
	return nil, nil
}
