package main

import (
	"context"
	"github.com/e1esm/LyrVibe/auth-service/api/v1/proto"
	"github.com/e1esm/LyrVibe/auth-service/internal/repository"
	"github.com/e1esm/LyrVibe/auth-service/internal/repository/sessionRepository/redis"
	"github.com/e1esm/LyrVibe/auth-service/internal/repository/userRepository/postgres"
	"github.com/e1esm/LyrVibe/auth-service/internal/server"
	"github.com/e1esm/LyrVibe/auth-service/internal/service"
	"github.com/e1esm/LyrVibe/auth-service/pkg/config"
	"github.com/e1esm/LyrVibe/auth-service/pkg/logger"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"net/http"
	"sync"
)

func main() {
	cfg := *config.NewConfig()
	authServer := configureServer(configureService(*configureRepositories(cfg), *service.NewTokenServiceBuilder()))

	listener, err := net.Listen("tcp", cfg.GRPC.Address)
	if err != nil {
		logger.Logger.Error(err.Error())
	}

	mux := configureHTTPServer(cfg.GRPC.Address)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		if err := http.ListenAndServe(cfg.HTTP.Address, mux); err != nil {
			logger.Logger.Error(err.Error())
		}
	}()
	go func() {
		defer wg.Done()
		if err := authServer.Server.Serve(listener); err != nil {
			logger.Logger.Error(err.Error())
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

func configureHTTPServer(address string) *runtime.ServeMux {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	err := proto.RegisterAuthServiceHandlerFromEndpoint(context.Background(), mux, address, opts)
	if err != nil {
		logger.Logger.Fatal("Couldn't have registered server from endpoint")
		return nil
	}
	return mux
}

func configureService(repositories repository.Repositories, manager service.TokenServiceBuilder) service.Service {
	return service.NewAuthService(repositories, manager)
}
