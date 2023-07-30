package redis

import (
	"fmt"
	"github.com/e1esm/LyrVibe/auth-service/pkg/config"
	"github.com/e1esm/LyrVibe/auth-service/pkg/logger"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"os"
)

type SessionStorage interface {
}

type SessionsRepository struct {
	redis *redis.Client
}

func NewSessionsStorage(config config.Config) SessionStorage {

	if err := godotenv.Load("redis.env"); err != nil {
		logger.Logger.Error(err.Error())
		return nil
	}

	password := os.Getenv("REDIS_PASSWORD")
	if password == "" {
		logger.Logger.Error("Password is not set")
		return nil
	}

	storageURL := fmt.Sprintf("%s:%d", config.SessionsStorage.ContainerName, config.SessionsStorage.Port)
	cli := redis.NewClient(&redis.Options{
		Addr:     storageURL,
		Password: password,
		DB:       0,
	})
	if cli == nil {
		logger.Logger.Error("Error")
	}
	return SessionsRepository{redis: cli}

}
