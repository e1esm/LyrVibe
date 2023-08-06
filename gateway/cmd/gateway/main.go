package main

import (
	"github.com/e1esm/LyrVibe/gateway/pkg/config"
	"github.com/e1esm/LyrVibe/gateway/pkg/logger"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	if err := godotenv.Load("config.yml"); err != nil {
		logger.Logger.Fatal("Couldn't have loaded config file",
			zap.String("err", err.Error()))
	}
	_ = config.NewConfig()

}
