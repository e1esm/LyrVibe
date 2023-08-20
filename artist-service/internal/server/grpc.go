package server

import (
	"context"
	"fmt"
	"github.com/e1esm/LyrVibe/artist-service/api/v1/proto"
	"github.com/e1esm/LyrVibe/artist-service/internal/service"
	"github.com/e1esm/LyrVibe/artist-service/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	validationError = "validation's failed: %s"
	verifyingError  = "couldn't validate artist with username: %s"
	verifiedUser    = "user with nickname: %s is successfully verified"
)

var artistRole = "Artist"

type Server struct {
	Server   *grpc.Server
	Services service.Services
	proto.UnimplementedArtistServiceServer
}

func (s *Server) Verify(ctx context.Context, request *proto.VerificationRequest) (*proto.VerificationResponse, error) {
	if err := request.ValidateAll(); err != nil {
		logger.Logger.Error(err.Error())
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf(validationError, err.Error()))
	}
	_, err := s.Services.ArtistService.AddArtist(ctx, request)
	if err != nil {
		logger.Logger.Error(err.Error())
		return nil, status.Error(codes.Internal, fmt.Sprintf(verifyingError, request.Username))
	}
	_, err = s.Services.RoleService.UpdateRole(ctx, request.Id, artistRole)
	if err != nil {
		return nil, err
	}
	return &proto.VerificationResponse{
		IsVerified:    true,
		RequestStatus: proto.RequestStatus_OK,
	}, nil
}
