package repository

import (
	"context"
	"fmt"
	"github.com/e1esm/LyrVibe/music-service/internal/entity"
	"github.com/e1esm/LyrVibe/music-service/pkg/config"
	"github.com/e1esm/LyrVibe/music-service/pkg/logger"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	NewTrack(entity.TrackEntity) (entity.TrackEntity, error)
}

type MusicRepository struct {
	Pool *pgxpool.Pool
}

func NewMusicRepository(cfg *config.Config) Repository {
	repo := MusicRepository{}
	var err error
	repo.Pool, err = pgxpool.New(context.Background(),
		fmt.Sprintf("%s://%s:%s@%s:%d/%s?pool_max_conns=%d",
			cfg.MusicStorage.Database,
			cfg.MusicStorage.User,
			cfg.MusicStorage.Password,
			cfg.MusicStorage.Address,
			cfg.MusicStorage.Port,
			cfg.MusicStorage.Database,
			cfg.MusicStorage.MaxConnections))
	if err != nil {
		logger.GetLogger().Error(err.Error())
		return nil
	}
	return &repo
}

func (mr *MusicRepository) NewTrack(track entity.TrackEntity) (entity.TrackEntity, error) {

	return track, nil
}
