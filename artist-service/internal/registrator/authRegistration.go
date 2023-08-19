package registrator

import (
	"fmt"
	"github.com/e1esm/LyrVibe/artist-service/pkg/config"
	"github.com/e1esm/LyrVibe/artist-service/pkg/logger"
	"github.com/e1esm/LyrVibe/auth-service/api/v1/proto"
	"google.golang.org/grpc"
)

func RegisterAuthService(config config.Config) proto.AuthServiceClient {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d",
		config.AuthService.ContainerName,
		config.AuthService.Port))
	if err != nil {
		logger.Logger.Error(err.Error())
		return nil
	}
	return proto.NewAuthServiceClient(conn)
}
