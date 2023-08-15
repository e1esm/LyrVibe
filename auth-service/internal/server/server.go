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
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	NoUserFound        = "User with username of %s wasn't found in the database or password was incorrect: %s"
	LogoutErr          = "Couldn't have logged out"
)

type Server struct {
	Server      *grpc.Server
	AuthService service.Service
	proto.UnimplementedAuthServiceServer
}

func (s *Server) SignUp(ctx context.Context, request *proto.SignUpRequest) (*proto.SignUpResponse, error) {
	err := request.ValidateAll()
	if err != nil {
		return &proto.SignUpResponse{
			Username: request.Username,
			Status: &proto.RequestStatus{
				RequestStatus: string(Fail),
				ErrorMessage:  err.Error(),
			},
		}, err
	}
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

	err = s.AuthService.SaveUser(ctx, user)
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
		},
	}, nil
}

func (s *Server) SignIn(ctx context.Context, request *proto.SignInRequest) (*proto.SignInResponse, error) {
	err := request.ValidateAll()
	if err != nil {
		return &proto.SignInResponse{
			Tokens: &proto.CachedTokens{},
			Status: &proto.RequestStatus{
				RequestStatus: string(Fail),
				ErrorMessage:  err.Error(),
			},
		}, err
	}
	user, err := s.AuthService.GetUser(ctx, request.Username, request.Password)
	if err == sql.ErrNoRows {
		return &proto.SignInResponse{
			Tokens: &proto.CachedTokens{},
			Status: &proto.RequestStatus{
				RequestStatus: string(Fail),
				ErrorMessage:  fmt.Sprintf(NoUserFound, request.Username, request.Password),
			},
		}, err
	}

	cachedTokens, err := s.AuthService.CreateSession(ctx, user)
	if err != nil {
		return &proto.SignInResponse{
			Tokens: &proto.CachedTokens{},
			Status: &proto.RequestStatus{
				RequestStatus: string(Fail),
				ErrorMessage:  err.Error(),
			},
		}, err
	}

	return &proto.SignInResponse{
		Tokens: &proto.CachedTokens{
			AccessToken:  cachedTokens.AccessToken,
			RefreshToken: cachedTokens.RefreshToken,
			AccessTTL:    fmt.Sprintf("%s", cachedTokens.AccessTTL),
			RefreshTTL:   fmt.Sprintf("%s", cachedTokens.RefreshTTL),
		},
		Status: &proto.RequestStatus{
			RequestStatus: string(Success),
		},
	}, nil
}

func (s *Server) Logout(ctx context.Context, request *proto.LogoutRequest) (*emptypb.Empty, error) {
	err := s.AuthService.Logout(ctx, request.AccessToken)
	if err != nil {
		return &emptypb.Empty{}, status.Error(codes.Internal, LogoutErr)
	}
	return &emptypb.Empty{}, nil
}

func (s *Server) UpdateRole(ctx context.Context, request *proto.UpdatingRoleRequest) (*proto.UpdatingRoleResponse, error) {
	id, err := uuid.FromBytes([]byte(request.UserId))
	if err != nil {
		return &proto.UpdatingRoleResponse{Status: BadRequestError}, status.Error(codes.InvalidArgument, BadRequestError)
	}
	err = s.AuthService.UpdateRole(ctx, id, models.Role(request.RequestedRole))
	if err != nil {
		return &proto.UpdatingRoleResponse{Status: InternalError}, status.Error(codes.Internal, InternalError)
	}
	return &proto.UpdatingRoleResponse{Status: "OK"}, nil
}
