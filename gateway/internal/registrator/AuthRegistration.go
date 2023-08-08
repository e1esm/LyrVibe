package registrator

import (
	"fmt"
	"github.com/e1esm/LyrVibe/auth-service/api/v1/proto"
	"github.com/e1esm/LyrVibe/gateway/pkg/config"
	"github.com/e1esm/LyrVibe/gateway/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RegisterAuthService(cfg config.Config) proto.AuthServiceClient {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", cfg.AuthService.Address,
		cfg.AuthService.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Logger.Error(err.Error())
		return nil
	}
	return proto.NewAuthServiceClient(conn)
}
