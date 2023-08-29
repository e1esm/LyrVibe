package registrator

import (
	"fmt"
	"github.com/e1esm/LyrVibe/artist-service/pkg/config"
	"github.com/e1esm/LyrVibe/artist-service/pkg/logger"
	"github.com/e1esm/LyrVibe/music-service/api/v1/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func RegisterMusicService(cfg *config.Config) proto.MusicServiceClient {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d",
		cfg.MusicServiceServer.ContainerName,
		cfg.MusicServiceServer.Port))
	if err != nil {
		logger.GetLogger().Fatal("Couldn't have established connection with Music Server",
			zap.String("err", err.Error()))
	}
	return proto.NewMusicServiceClient(conn)
}
