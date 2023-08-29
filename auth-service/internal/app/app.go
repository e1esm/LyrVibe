package app

import (
	"github.com/e1esm/LyrVibe/auth-service/api/v1/proto"
	"github.com/e1esm/LyrVibe/auth-service/internal/repository"
	"github.com/e1esm/LyrVibe/auth-service/internal/repository/sessionRepository/redis"
	"github.com/e1esm/LyrVibe/auth-service/internal/repository/userRepository/postgres"
	"github.com/e1esm/LyrVibe/auth-service/internal/server"
	"github.com/e1esm/LyrVibe/auth-service/internal/service"
	"github.com/e1esm/LyrVibe/auth-service/pkg/config"
	"github.com/e1esm/LyrVibe/auth-service/pkg/logger"
	"google.golang.org/grpc"
	"net"
	"sync"
)

func Run() {
	cfg := *config.NewConfig()
	authServer := configureServer(configureService(*configureRepositories(cfg), *service.NewTokenServiceBuilder()))

	listener, err := net.Listen("tcp", cfg.GRPC.Address)
	if err != nil {
		logger.GetLogger().Error(err.Error())
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := authServer.Server.Serve(listener); err != nil {
			logger.GetLogger().Error(err.Error())
		}
	}()

	wg.Wait()
}

func configureRepositories(config config.Config) *repository.Repositories {
	return repository.NewRepositoriesBuilder().
		WithMainRepo(postgres.NewUserRepository(config)).
		WithSessionsStorage(redis.NewSessionsStorage(config)).
		Build()
}

func configureServer(authService service.Service) *server.Server {
	authServer := server.Server{}
	authServer.Server = grpc.NewServer([]grpc.ServerOption{}...)
	authServer.AuthService = authService
	proto.RegisterAuthServiceServer(authServer.Server, &authServer)

	return &authServer
}

func configureService(repositories repository.Repositories, manager service.TokenServiceBuilder) service.Service {
	return service.NewAuthService(repositories, manager)
}
