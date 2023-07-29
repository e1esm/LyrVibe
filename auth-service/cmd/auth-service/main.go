package main

import (
	"github.com/e1esm/LyrVibe/auth-service/internal/repository/sessionRepository/redis"
	"github.com/e1esm/LyrVibe/auth-service/internal/repository/userRepository/postgres"
	"github.com/e1esm/LyrVibe/auth-service/pkg/config"
	"time"
)

func main() {
	_ = postgres.NewUserRepository(*config.NewConfig())
	_ = redis.NewSessionsStorage(*config.NewConfig())
	time.Sleep(100 * time.Second)
}
