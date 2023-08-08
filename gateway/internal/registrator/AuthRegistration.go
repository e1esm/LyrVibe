package registrator

import (
	"fmt"
	"github.com/e1esm/LyrVibe/auth-service/api/v1/proto"
	"github.com/e1esm/LyrVibe/gateway/pkg/config"
	"github.com/e1esm/LyrVibe/gateway/pkg/logger"
	"google.golang.org/grpc"
)

func RegisterAuthService(cfg config.Config) proto.AuthServiceClient {
	var opts []grpc.DialOption
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", cfg.AuthService.Address,
		cfg.AuthService.Port), opts...)
	if err != nil {
		logger.Logger.Error(err.Error())
		return nil
	}
	return proto.NewAuthServiceClient(conn)
}
