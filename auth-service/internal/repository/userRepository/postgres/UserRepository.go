package postgres

import (
	"context"
	"fmt"
	"github.com/e1esm/LyrVibe/auth-service/pkg/config"
	"github.com/e1esm/LyrVibe/auth-service/pkg/logger"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"os"
)

type UserStorage interface {
}

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(config config.Config) UserStorage {

	if err := godotenv.Load("postgres.env"); err != nil {
		logger.Logger.Error(err.Error())
		return nil
	}
	pgUser := os.Getenv("POSTGRES_USER")
	pgPassword := os.Getenv("POSTGRES_PASSWORD")
	pgDB := os.Getenv("POSTGRES_DB")
	pgPort := os.Getenv("POSTGRES_PORT")

	databaseUrl := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?pool_max_conns=%d",
		pgUser,
		pgPassword,
		config.UsersStorage.ContainerName,
		pgPort,
		pgDB,
		config.UsersStorage.MaxConnectionPool)
	pool, err := pgxpool.New(context.Background(), databaseUrl)
	if err != nil {
		logger.Logger.Error(err.Error(), zap.String("url", databaseUrl))
		return nil
	}

	return &UserRepository{pool: pool}
}
