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
	artist, err := s.Services.ArtistService.AddArtist(ctx, request)
	if err != nil {
		logger.Logger.Error(err.Error())
		return nil, status.Error(codes.Internal, fmt.Sprintf(verifyingError, request.Username))
	}
	return &proto.VerificationResponse{
		IsVerified: true,
		RequestStatus: &proto.RequestStatus{
			RequestStatus: fmt.Sprintf(verifiedUser, artist.Username),
			ErrorMessage:  "",
		},
	}, nil
}
