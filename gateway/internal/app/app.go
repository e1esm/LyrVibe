package app

import (
	"fmt"
	"github.com/e1esm/LyrVibe/gateway/internal/server"
	"github.com/e1esm/LyrVibe/gateway/internal/service"
	"github.com/e1esm/LyrVibe/gateway/pkg/config"
	"github.com/e1esm/LyrVibe/gateway/pkg/logger"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"sync"
)

func Run() {
	if err := godotenv.Load("config.yml"); err != nil {
		logger.Logger.Fatal("Couldn't have loaded config file",
			zap.String("err", err.Error()))
	}
	cfg := *config.NewConfig()
	proxy := setUpServer(configureServices(cfg))
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		proxy.Run(fmt.Sprintf("%s:%s", cfg.Server.Address, cfg.Server.Port))
	}()
	wg.Wait()
}

func configureServices(cfg config.Config) service.Services {
	return service.NewServiceBuilder().
		WithAuthProvider(setUpAuthService(cfg)).
		Build()
}

func setUpServer(services service.Services) server.Proxy {
	return server.NewProxyServer(services)
}

func setUpAuthService(cfg config.Config) service.AuthenticationProvider {
	return service.NewAuthenticationService(cfg)
}
