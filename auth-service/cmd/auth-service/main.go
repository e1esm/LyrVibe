package main

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
)

func main() {
	cfg := *config.NewConfig()
	authServer := configureServer(configureService(*configureRepositories(cfg), service.NewTokenServiceBuilder()))

	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		logger.Logger.Error(err.Error())
	}

	if err = authServer.Server.Serve(listener); err != nil {
		logger.Logger.Error(err.Error())
	}
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
