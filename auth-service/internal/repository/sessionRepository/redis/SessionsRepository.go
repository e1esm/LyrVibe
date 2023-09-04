package redis

import (
	"context"
	"errors"
	"fmt"
	"github.com/e1esm/LyrVibe/auth-service/internal/models"
	"github.com/e1esm/LyrVibe/auth-service/pkg/config"
	"github.com/e1esm/LyrVibe/auth-service/pkg/logger"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"os"
)

var (
	expiredErr = errors.New("session's expired or have never been started")
)

type SessionStorage interface {
	Add(context.Context, models.CachedTokens) (bool, error)
	Get(context.Context, string) (models.CachedTokens, error)
	Delete(context.Context, string) error
}

type SessionsRepository struct {
	redis *redis.Client
}

func NewSessionsStorage(config config.Config) SessionStorage {

	if err := godotenv.Load("redis.env"); err != nil {
		logger.GetLogger().Error(err.Error())
		return nil
	}

	password := os.Getenv("REDIS_PASSWORD")
	if password == "" {
		logger.GetLogger().Error("Password is not set")
		return nil
	}

	storageURL := fmt.Sprintf("%s:%d", config.SessionsStorage.ContainerName, config.SessionsStorage.Port)
	cli := redis.NewClient(&redis.Options{
		Addr:     storageURL,
		Password: password,
		DB:       0,
	})
	if cli == nil {
		logger.GetLogger().Error("Error")
	}
	return &SessionsRepository{redis: cli}

}

func (sr *SessionsRepository) Get(ctx context.Context, refreshToken string) (models.CachedTokens, error) {
	cmd := sr.redis.Get(ctx, fmt.Sprintf("%v", refreshToken))
	var cachedTokens models.CachedTokens
	if err := cmd.Scan(&cachedTokens); err != nil {
		logger.GetLogger().Error("Error while getting credentials from Redis",
			zap.String("err", err.Error()))
		return models.CachedTokens{}, expiredErr
	}
	return cachedTokens, nil
}

func (sr *SessionsRepository) Add(ctx context.Context, tokens models.CachedTokens) (bool, error) {
	isOk, err := sr.redis.SetNX(ctx, fmt.Sprintf("%v", tokens.RefreshToken), tokens, tokens.RefreshTTL).Result()
	return isOk, err
}

func (sr *SessionsRepository) Delete(ctx context.Context, id string) error {
	res := sr.redis.Del(ctx, id)
	return res.Err()
}
