package server

import (
	"github.com/e1esm/LyrVibe/music-service/api/v1/proto"
	"google.golang.org/grpc"
)

type Server struct {
	server *grpc.Server
	proto.UnimplementedMusicServiceServer
	Services Services
}
