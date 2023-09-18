package server

import (
	"context"
	"fmt"
	artist "github.com/e1esm/LyrVibe/artist-service/api/v1/proto"
	"github.com/e1esm/LyrVibe/artist-service/internal/models"
	"github.com/e1esm/LyrVibe/artist-service/internal/service"
	"github.com/e1esm/LyrVibe/artist-service/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	validationError = "validation's failed: %s"
	verifyingError  = "couldn't validate artist with username: %s"
	deleteErr       = "can't delete the track: %v"
)

var artistRole = "Artist"

type Server struct {
	Server   *grpc.Server
	Services service.Services
	artist.UnimplementedArtistServiceServer
}

func (s *Server) Verify(ctx context.Context, request *artist.VerificationRequest) (*artist.VerificationResponse, error) {
	if err := request.ValidateAll(); err != nil {
		logger.GetLogger().Error(err.Error())
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf(validationError, err.Error()))
	}
	_, err := s.Services.ArtistService.AddArtist(ctx, request)
	if err != nil {
		logger.GetLogger().Error(err.Error())
		return nil, status.Error(codes.Internal, fmt.Sprintf(verifyingError, request.Username))
	}
	_, err = s.Services.RoleService.UpdateRole(ctx, request.Id, artistRole)
	if err != nil {
		return nil, err
	}
	return &artist.VerificationResponse{
		IsVerified:    true,
		RequestStatus: artist.RequestStatus_OK,
	}, nil
}

func (s *Server) AddTrack(ctx context.Context, request *artist.NewTrackRequest) (*artist.NewTrackResponse, error) {
	resp, err := s.Services.MusicService.Release(ctx, models.NewSong(request))
	if err != nil {
		return nil, fmt.Errorf("can't release track: %v", err)
	}
	return &artist.NewTrackResponse{
		Title:         resp.Title,
		RequestStatus: artist.RequestStatus_OK,
	}, nil
}

func (s *Server) DeleteTrack(ctx context.Context, req *artist.DeleteTrackRequest) (*artist.DeleteTrackResponse, error) {
	resp, err := s.Services.MusicService.DeleteTrack(ctx, req)
	if err != nil {
		return nil, fmt.Errorf(deleteErr, err.Error())
	}
	return resp, nil
}
