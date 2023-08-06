package main

import (
	"context"
	"fmt"
	"github.com/e1esm/LyrVibe/gateway/internal/interceptors"
	"github.com/e1esm/LyrVibe/gateway/internal/registers"
	"github.com/e1esm/LyrVibe/gateway/pkg/config"
	"github.com/e1esm/LyrVibe/gateway/pkg/logger"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/joho/godotenv"
	"github.com/tmc/grpc-websocket-proxy/wsproxy"
	"go.uber.org/zap"
	"net/http"
)

func main() {
	if err := godotenv.Load("config.yml"); err != nil {
		logger.Logger.Fatal("Couldn't have loaded config file",
			zap.String("err", err.Error()))
	}
	cfg := config.NewConfig()
	_, err := interceptors.NewAuthInterceptor(*cfg)
	if err != nil {
		logger.Logger.Fatal("Couldn't have set up auth interceptor")
	}
	mux := runtime.NewServeMux()
	err = registers.RegisterAuth(context.Background(), mux,
		fmt.Sprintf("%s:%s", cfg.AuthService.Address, cfg.AuthService.Port))
	if err != nil {
		logger.Logger.Error(err.Error())
	}
	if err = http.ListenAndServe("", mux); err != nil {
		logger.Logger.Fatal("Couldn't have started server", zap.String("err", err.Error()))

	}
	ochttpHandler, err := customHandler(mux)
	if err != nil {
		logger.Logger.Error(err.Error())
	}

	server := newGatewayServer(":8080", ochttpHandler)

	err = server.ListenAndServe()
	if err != nil {
		logger.Logger.Error(err.Error())
	}
}

func newGatewayServer(addr string, oc http.Handler) *http.Server {
	return &http.Server{
		Addr:    addr,
		Handler: oc,
	}
}

func customHandler(mux *runtime.ServeMux) (http.Handler, error) {
	if err := mux.HandlePath(
		"GET",
		"/health",
		func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
			if _, err := w.Write([]byte("Gateway v1 serving")); err != nil {
				logger.Logger.Error("Health check error", zap.String("err", err.Error()))
			}
		},
	); err != nil {
		return nil, err
	}

	return wsproxy.WebsocketProxy(mux), nil
}
