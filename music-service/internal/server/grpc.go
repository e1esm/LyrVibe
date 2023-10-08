package server

import (
	"context"
	"fmt"
	"github.com/e1esm/LyrVibe/music-service/api/v1/proto"
	"github.com/e1esm/LyrVibe/music-service/internal/service"
	"github.com/e1esm/LyrVibe/music-service/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	newTrackErr = "track wasn't added"
)

type Server struct {
	Server *grpc.Server
	proto.UnimplementedMusicServiceServer
	Services service.Services
}

func NewServer(server *grpc.Server, services service.Services) *Server {
	return &Server{
		Server:   server,
		Services: services,
	}
}

func (s *Server) AddNewTrack(ctx context.Context, request *proto.NewTrackRequest) (*proto.NewTrackResponse, error) {
	if err := request.ValidateAll(); err != nil {
		return nil, fmt.Errorf("validation error: %v", err)
	}
	track, err := s.Services.MusicService.AddNew(ctx, request)
	if err != nil {
		logger.GetLogger().Error(err.Error())
		return nil, status.Error(codes.Internal, newTrackErr)
	}
	return &proto.NewTrackResponse{
		Status: proto.Status_OK,
		Title:  track.Data.Title,
	}, nil
}

func (s *Server) DeleteTrack(ctx context.Context, request *proto.DeleteRequest) (*proto.DeleteResponse, error) {

	if err := request.ValidateAll(); err != nil {
		return nil, fmt.Errorf("validation error: %v", err)
	}

	resp, err := s.Services.MusicService.Delete(ctx, request)
	if err != nil {
		logger.GetLogger().Error(err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}
