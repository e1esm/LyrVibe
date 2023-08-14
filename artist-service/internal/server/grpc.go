package server

import (
	"github.com/e1esm/LyrVibe/artist-service/api/v1/proto"
	"github.com/e1esm/LyrVibe/artist-service/internal/service"
	"google.golang.org/grpc"
)

type Server struct {
	Server   *grpc.Server
	Services service.Services
	proto.UnimplementedArtistServiceServer
}
