package server

import "github.com/e1esm/LyrVibe/artist-service/api/v1/proto"

type Server struct {
	Server proto.ArtistServiceServer
}

func NewServer() *Server {

}
