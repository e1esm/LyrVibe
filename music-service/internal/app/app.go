package app

import (
	"github.com/e1esm/LyrVibe/music-service/api/v1/proto"
	"github.com/e1esm/LyrVibe/music-service/internal/repository"
	"github.com/e1esm/LyrVibe/music-service/internal/server"
	"github.com/e1esm/LyrVibe/music-service/internal/service"
	"github.com/e1esm/LyrVibe/music-service/pkg/config"
	"github.com/e1esm/LyrVibe/music-service/pkg/logger"
	"google.golang.org/grpc"
	"net"
	"sync"
)

func Run() {
	cfg := config.NewConfig()
	listener, err := net.Listen("tcp", cfg.GRPC.Address)
	if err != nil {
		logger.GetLogger().Fatal(err.Error())
	}
	musicServer := setUpServer(cfg)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := musicServer.Server.Serve(listener); err != nil {
			logger.GetLogger().Fatal(err.Error())
		}
	}()
	wg.Wait()
}

func setUpServer(cfg *config.Config) *server.Server {
	grpcServer := setUpGRPCConnection(cfg)
	musicServer := server.NewServer(grpcServer, setUpServices(cfg))
	proto.RegisterMusicServiceServer(grpcServer, musicServer)
	return musicServer
}

func setUpServices(cfg *config.Config) service.Services {
	return service.NewServicesBuilder().
		WithMusicService(service.NewMusicService(
			repository.NewMusicRepository(cfg),
		),
		).Build()
}

func setUpGRPCConnection(cfg *config.Config) *grpc.Server {
	return grpc.NewServer(grpc.EmptyServerOption{})
}
