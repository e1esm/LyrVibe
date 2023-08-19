package repository

import (
	"context"
	"fmt"
	"github.com/e1esm/LyrVibe/artist-service/internal/models"
	"github.com/e1esm/LyrVibe/artist-service/pkg/config"
	"github.com/e1esm/LyrVibe/artist-service/pkg/logger"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"time"
)

const (
	timeForReq = time.Second * 5
)

type Repository interface {
	Add(context.Context, *models.Artist) error
}

type ArtistRepository struct {
	pool *pgxpool.Pool
}

func NewRepository(cfg *config.Config) Repository {
	dsn := fmt.Sprintf("%s://%s:%s@%s:%d/%s?pool_max_conns=%d",
		cfg.ArtistStorage.Database,
		cfg.ArtistStorage.DatabaseUser,
		cfg.ArtistStorage.DatabasePassword,
		cfg.ArtistStorage.ContainerName,
		cfg.ArtistStorage.Port,
		cfg.ArtistStorage.DatabaseName,
		cfg.ArtistStorage.MaxConnections)
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		logger.Logger.Error(err.Error())
		return nil
	}
	return &ArtistRepository{pool: pool}
}

func (ar *ArtistRepository) Add(ctx context.Context, artist *models.Artist) error {
	ctx, cancel := context.WithTimeout(ctx, timeForReq)
	defer cancel()
	_, err := ar.pool.Exec(ctx, "INSERT INTO artists VALUES ($1, $2, $3, $4, $5, $6)",
		artist.ID,
		artist.Username,
		artist.Country,
		artist.FirstName,
		artist.SecondName,
		artist.Views)
	if err != nil {
		logger.Logger.Error("Error while operating over request", zap.String("err", err.Error()))
		return err
	}
	return nil
}
