package app

import (
	"github.com/e1esm/LyrVibe/artist-service/api/v1/proto"
	"github.com/e1esm/LyrVibe/artist-service/internal/registrator"
	"github.com/e1esm/LyrVibe/artist-service/internal/repository"
	"github.com/e1esm/LyrVibe/artist-service/internal/server"
	"github.com/e1esm/LyrVibe/artist-service/internal/service"
	"github.com/e1esm/LyrVibe/artist-service/pkg/config"
	"github.com/e1esm/LyrVibe/artist-service/pkg/logger"
	"google.golang.org/grpc"
	"net"
	"sync"
)

func Run() {
	cfg := config.NewConfig()
	listener, err := net.Listen("tcp", cfg.Server.Address)
	if err != nil {
		logger.GetLogger().Fatal(err.Error())
	}
	artistServer := setupServer(setupServices(setupRepository(cfg), cfg))
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		if err = artistServer.Server.Serve(listener); err != nil {
			logger.GetLogger().Fatal(err.Error())
		}
	}()
	logger.GetLogger().Info("Server's started")
	wg.Wait()
}

func setupServer(services service.Services) *server.Server {
	artistServer := server.Server{}
	artistServer.Server = grpc.NewServer(grpc.EmptyServerOption{})
	artistServer.Services = services
	proto.RegisterArtistServiceServer(artistServer.Server, &artistServer)
	return &artistServer
}

func setupServices(repo repository.Repository, cfg *config.Config) service.Services {
	return service.NewServiceBuilder().
		WithArtistService(service.NewArtistService(repo)).
		WithRoleService(service.NewRolesService(registrator.RegisterRoleService(cfg))).
		WithMusicService(service.NewMusicService(registrator.RegisterMusicService(cfg))).
		Build()
}

func setupRepository(cfg *config.Config) repository.Repository {
	return repository.NewRepository(cfg)
}
