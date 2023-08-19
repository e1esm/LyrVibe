package registrator

import (
	"fmt"
	"github.com/e1esm/LyrVibe/artist-service/pkg/config"
	"github.com/e1esm/LyrVibe/artist-service/pkg/logger"
	"github.com/e1esm/LyrVibe/auth-service/api/v1/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RegisterRoleService(config config.Config) proto.AuthServiceClient {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d",
		config.AuthService.ContainerName,
		config.AuthService.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Logger.Error(err.Error())
		return nil
	}
	return proto.NewAuthServiceClient(conn)
}
