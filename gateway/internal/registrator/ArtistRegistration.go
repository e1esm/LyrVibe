package registrator

import (
	"fmt"
	"github.com/e1esm/LyrVibe/artist-service/api/v1/proto"
	"github.com/e1esm/LyrVibe/gateway/pkg/config"
	"github.com/e1esm/LyrVibe/gateway/pkg/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RegisterArtistService(cfg *config.Config) proto.ArtistServiceClient {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", cfg.ArtistService.Address, cfg.ArtistService.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Logger.Error("Error while establishing connection with ArtistService", zap.String("err", err.Error()))
		return nil
	}
	return proto.NewArtistServiceClient(conn)
}
