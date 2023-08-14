package app

import (
	"github.com/e1esm/LyrVibe/artist-service/api/v1/proto"
	"github.com/e1esm/LyrVibe/artist-service/internal/repository"
	"github.com/e1esm/LyrVibe/artist-service/internal/server"
	"github.com/e1esm/LyrVibe/artist-service/internal/service"
	"github.com/e1esm/LyrVibe/artist-service/pkg/config"
	"github.com/e1esm/LyrVibe/artist-service/pkg/logger"
	"google.golang.org/grpc"
	"net"
)

func Run() {
	cfg := config.NewConfig()
	listener, err := net.Listen("tcp", cfg.Server.Address)
	if err != nil {
		logger.Logger.Fatal(err.Error())
	}
	artistServer := setupServer(setupServices(setupRepository(cfg)))
	if err = artistServer.Server.Serve(listener); err != nil {
		logger.Logger.Fatal(err.Error())
	}
}

func setupServer(services service.Services) *server.Server {
	artistServer := server.Server{}
	artistServer.Server = grpc.NewServer(grpc.EmptyServerOption{})
	proto.RegisterArtistServiceServer(artistServer.Server, &artistServer)
	return &artistServer
}

func setupServices(repo repository.Repository) service.Services {
	return service.NewServiceBuilder().
		WithArtistService(service.NewArtistService(repo)).
		Build()
}

func setupRepository(cfg *config.Config) repository.Repository {
	return repository.NewRepository(cfg)
}
