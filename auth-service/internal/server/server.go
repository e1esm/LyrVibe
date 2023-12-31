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
	"go.uber.org/zap"
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
	BadRequestError = "bad request"
	InternalError   = "internal error occurred while operating on the provided input"
	NoUserFound     = "User with username of %s wasn't found in the database or password was incorrect: %s"
	LogoutErr       = "Couldn't have logged out"
	SessionErr      = "No session found"
	UpdatingErr     = "Error while updating session"
)

type Server struct {
	Server      *grpc.Server
	AuthService service.Service
	proto.UnimplementedAuthServiceServer
}

func (s *Server) SignUp(ctx context.Context, request *proto.SignUpRequest) (*proto.SignUpResponse, error) {
	err := request.ValidateAll()
	if err != nil {
		return nil, err
	}
	user := models.NewUser(request)
	if user == nil {
		return nil, status.Error(codes.InvalidArgument, BadRequestError)
	}

	err = s.AuthService.SaveUser(ctx, user)
	if err != nil {
		logger.GetLogger().Error(err.Error())
		return nil, status.Error(codes.Internal, InternalError)
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
		logger.GetLogger().Error("Validation error", zap.String("err", err.Error()))
		return nil, err
	}
	user, err := s.AuthService.GetUser(ctx, request.Username, request.Password)
	if errors.Is(err, sql.ErrNoRows) {
		logger.GetLogger().Error("No found err", zap.String("err", err.Error()))
		return nil, status.Error(codes.NotFound, fmt.Sprintf(NoUserFound, request.Username, request.Password))
	}

	cachedTokens, err := s.AuthService.CreateSession(ctx, user)
	if err != nil {
		logger.GetLogger().Error("Sessions not created", zap.String("err", err.Error()))
		return nil, err
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
	err := s.AuthService.Logout(ctx, request.RefreshToken)
	if err != nil {
		logger.GetLogger().Error("Logout error", zap.String("err", err.Error()))
		return nil, status.Error(codes.Internal, LogoutErr)
	}
	return &emptypb.Empty{}, nil
}

func (s *Server) UpdateRole(ctx context.Context, request *proto.UpdatingRoleRequest) (*proto.UpdatingRoleResponse, error) {
	id := uuid.MustParse(request.UserId)
	err := s.AuthService.UpdateRole(ctx, id, models.Role(request.RequestedRole))
	if err != nil {
		return nil, status.Error(codes.Internal, InternalError)
	}
	return &proto.UpdatingRoleResponse{Status: "OK"}, nil
}

func (s *Server) Verification(ctx context.Context, request *proto.VerificationRequest) (*proto.VerificationResponse, error) {
	payload, err := s.AuthService.GetPayload(request.AccessToken)
	if err != nil {
		logger.GetLogger().Error("Couldn't have gotten credentials", zap.String("err", err.Error()))
		return nil, status.Error(codes.Internal, InternalError)
	}
	return &proto.VerificationResponse{Role: string(payload.Role), Id: payload.ID, Username: payload.Username}, nil
}

func (s *Server) RefreshToken(ctx context.Context, request *proto.RefreshRequest) (*proto.RefreshResponse, error) {
	tokens, err := s.AuthService.GetSessionCredentials(ctx, request.RefreshToken)
	if err != nil {
		logger.GetLogger().Error(err.Error())
		return nil, status.Error(codes.Internal, SessionErr)
	}

	payload, err := s.AuthService.GetPayload(tokens.AccessToken)
	id, err := uuid.Parse(payload.ID)
	if err != nil {
		logger.GetLogger().Error(err.Error())
		return nil, status.Error(codes.Internal, "Failed to parse UUID")
	}
	tokens, err = s.AuthService.CreateSession(ctx, &models.User{Username: payload.Username, ID: id, Role: payload.Role})
	logger.GetLogger().Info("Refresh Token", zap.String("", tokens.RefreshToken))
	_, err = s.AuthService.UpdateSession(ctx, tokens)
	if err != nil {
		logger.GetLogger().Error(err.Error())
		return nil, status.Error(codes.Internal, UpdatingErr)
	}

	return &proto.RefreshResponse{
		AccessToken: tokens.AccessToken,
		Ttl:         fmt.Sprintf("%v", tokens.AccessTTL),
	}, nil
}
