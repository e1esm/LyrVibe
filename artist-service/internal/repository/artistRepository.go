package repository

import (
	"context"
	"fmt"
	"github.com/e1esm/LyrVibe/artist-service/pkg/config"
	"github.com/e1esm/LyrVibe/artist-service/pkg/logger"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
}

type ArtistRepository struct {
	pool *pgxpool.Pool
}

func NewRepository(cfg *config.Config) Repository {
	dsn := fmt.Sprintf("%s://%s:%s@%s:%d/%s",
		cfg.ArtistStorage.Database,
		cfg.ArtistStorage.DatabaseUser,
		cfg.ArtistStorage.DatabasePassword,
		cfg.ArtistStorage.ContainerName,
		cfg.ArtistStorage.Port,
		cfg.ArtistStorage.DatabaseName)
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		logger.Logger.Error(err.Error())
		return nil
	}
	return &ArtistRepository{pool: pool}
}
