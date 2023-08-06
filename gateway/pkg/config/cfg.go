package config

import (
	"github.com/e1esm/LyrVibe/gateway/pkg/logger"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	AuthService struct {
		Address string `yaml:"addr"`
		Port    string `yaml:"port"`
	} `yaml:"auth_service"`
}

func NewConfig() *Config {
	bytes, err := os.ReadFile("config.yml")
	if err != nil {
		logger.Logger.Fatal("Couldn't have read config file")
		return nil
	}
	var config Config
	err = yaml.Unmarshal(bytes, &config)
	if err != nil {
		logger.Logger.Fatal("Couldn't have unmarshalled config file",
			zap.String("err", err.Error()))
		return nil
	}
	return &config
}
