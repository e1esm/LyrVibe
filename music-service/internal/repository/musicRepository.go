package repository

import (
	"context"
	"fmt"
	"github.com/e1esm/LyrVibe/music-service/internal/entity"
	"github.com/e1esm/LyrVibe/music-service/pkg/config"
	"github.com/e1esm/LyrVibe/music-service/pkg/logger"
	"github.com/e1esm/LyrVibe/music-service/pkg/uuidParser"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"time"
)

const (
	timeout             = 5 * time.Second
	maxConnParameterKey = "pool_max_conns"
	migrationsDir       = "file://db/migrations"
	sslmodeParameter    = "sslmode=disable"
)

type Repository interface {
	NewTrack(context.Context, entity.TrackEntity) (entity.TrackEntity, error)
}

type MusicRepository struct {
	pool            *pgxpool.Pool
	transactionRepo TransactionManager
}

func NewMusicRepository(cfg *config.Config, tManager TransactionManager) Repository {
	repo := MusicRepository{transactionRepo: tManager}
	var err error
	basicDBUrl := fmt.Sprintf("%s://%s:%s@%s:%d/%s",
		cfg.MusicStorage.Database,
		cfg.MusicStorage.User,
		cfg.MusicStorage.Password,
		cfg.MusicStorage.Address,
		cfg.MusicStorage.Port,
		cfg.MusicStorage.Name)

	repo.pool, err = connect(context.Background(),
		fmt.Sprintf("%s?%s=%d",
			basicDBUrl,
			maxConnParameterKey,
			cfg.MusicStorage.MaxConnections,
		))
	if err != nil {
		panic(err)
	}
	runMigrations(fmt.Sprintf("%s?%s", basicDBUrl, sslmodeParameter), migrationsDir, UP)
	return &repo
}

func (mr *MusicRepository) NewTrack(ctx context.Context, track entity.TrackEntity) (entity.TrackEntity, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	tx, err := mr.pool.Begin(ctx)
	if err != nil {
		logger.GetLogger().
			Error("Error while starting new transaction", zap.String("err", err.Error()))
		return entity.TrackEntity{}, err
	}
	mr.transactionRepo.AddTx(track.ID, tx)
	defer func() {
		mr.transactionRepo.Delete(track.ID)
		_ = tx.Rollback(ctx)
	}()
	err = mr.addTrack(ctx, track)
	if err != nil {
		return entity.TrackEntity{}, err
	}
	return track, nil
}

func (mr *MusicRepository) addTrack(ctx context.Context, track entity.TrackEntity) error {
	tx, err := mr.transactionRepo.Get(track.ID)
	if err != nil {
		return err
	}
	query := `INSERT INTO tracks
VALUES($1, $2, $3, $4, $6, $7, $8, $9, $10, $11, $12)`
	_, err = tx.Exec(ctx, query,
		track.ID,
		uuidParser.ParseUUID(track.Data.ArtistId),
		track.Data.Cover,
		track.Data.Title,
		track.Data.ReleaseDate,
		track.Data.Genre,
		track.Data.Genre,
		track.Data.Duration,
		track.Data.Country,
		track.Data.VideoLink,
		track.Data.Feature,
		track.CreatedAt,
	)
	if err != nil {
		logger.GetLogger().Error("Error while inserting into tracks table",
			zap.String("err", err.Error()))
		return err
	}
	return mr.addLyrics(ctx, track)
}

func (mr *MusicRepository) addLyrics(ctx context.Context, track entity.TrackEntity) error {
	batch := &pgx.Batch{}
	for i := 0; i < len(track.Data.Lyrics); i++ {
		batch.Queue("INSERT INTO lyrics (song_id, line, line_n) VALUES($1, $2, $3)",
			track.ID, track.Data.Lyrics[i], i+1)
	}
	tx, err := mr.transactionRepo.Get(track.ID)
	if err != nil {
		return err
	}
	err = tx.SendBatch(ctx, batch).Close()
	if err != nil {
		logger.GetLogger().Error("Couldn't perform a batch",
			zap.String("err", err.Error()))
		return err
	}
	return mr.addStatistics(ctx, track)
}

func (mr *MusicRepository) addStatistics(ctx context.Context, track entity.TrackEntity) error {
	tx, err := mr.transactionRepo.Get(track.ID)
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, "INSERT INTO track_statistics (song_id) VALUES ($1)", track.ID)
	return err
}
