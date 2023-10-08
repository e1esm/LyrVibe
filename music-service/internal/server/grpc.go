package server

import (
	"context"
	"fmt"
	"github.com/e1esm/LyrVibe/music-service/api/v1/proto"
	"github.com/e1esm/LyrVibe/music-service/internal/service"
	"github.com/e1esm/LyrVibe/music-service/pkg/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	newTrackErr   = "track wasn't added"
	validationErr = "validation error: %v"
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
		return nil, fmt.Errorf(validationErr, err)
	}
	track, err := s.Services.MusicService.AddNewTrack(ctx, request)
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
		return nil, fmt.Errorf(validationErr, err)
	}

	resp, err := s.Services.MusicService.Delete(ctx, request)
	if err != nil {
		logger.GetLogger().Error(err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) AddNewAlbum(ctx context.Context, request *proto.NewAlbumRequest) (*proto.NewAlbumResponse, error) {
	if err := request.ValidateAll(); err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Errorf(validationErr, err).Error())
	}
	resp, err := s.Services.MusicService.AddNewAlbum(ctx, request)
	if err != nil {
		logger.GetLogger().Error(err.Error(),
			zap.String("method", "NewAlbum"),
			zap.String("Title", request.Title))
		return nil, status.Error(codes.Internal, fmt.Errorf("creating error: %v", err).Error())
	}
	return &proto.NewAlbumResponse{
		Title:  resp.Title,
		Status: proto.Status_OK,
	}, nil
}
