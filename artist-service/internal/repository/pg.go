package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/e1esm/LyrVibe/artist-service/pkg/logger"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"time"
)

type MigrationType int

const (
	maxRetries               = 10
	UP         MigrationType = iota
	Down
)

func connect(ctx context.Context, dbURl string) (*pgxpool.Pool, error) {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()
	var pool *pgxpool.Pool
	var err error
	retry := 0
	for retry < maxRetries {
		select {
		case <-ticker.C:
			pool, err = pgxpool.New(ctx, dbURl)
			retry++
			if err := pool.Ping(context.Background()); err == nil {
				return pool, nil
			}
		}
	}
	return nil, fmt.Errorf("coudln't have connected to the database: %v", err)
}

func runMigrations(dbURL, fileDir string, migrationType MigrationType) {
	m, err := migrate.New(fileDir, dbURL)
	if err != nil {
		logger.GetLogger().Fatal("Couldn't have created migration",
			zap.String("err", err.Error()))
	}
	switch migrationType {
	case UP:
		if err = m.Up(); err != nil {
			if errors.Is(err, migrate.ErrNoChange) {
				logger.GetLogger().Error("Nothing to be changed")
				return
			}
			logger.GetLogger().Fatal("Couldn't have performed UP migration",
				zap.String("err", err.Error()))
		}
	case Down:
		if err = m.Down(); err != nil {
			if errors.Is(err, migrate.ErrNoChange) {
				logger.GetLogger().Error("Nothing to be changed")
				return
			}
			logger.GetLogger().Fatal("Couldn't have performed DOWN migration",
				zap.String("err", err.Error()))
		}
	}
}
