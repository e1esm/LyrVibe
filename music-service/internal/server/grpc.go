package server

import (
	"github.com/e1esm/LyrVibe/music-service/api/v1/proto"
	"github.com/e1esm/LyrVibe/music-service/internal/service"
	"google.golang.org/grpc"
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
