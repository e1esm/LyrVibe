package config

import (
	"github.com/e1esm/LyrVibe/gateway/pkg/logger"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	ArtistService struct {
		Address string `yaml:"addr"`
		Port    int    `yaml:"port"`
	} `yaml:"artist_service"`
	AuthService struct {
		Address string `yaml:"addr"`
		Port    int    `yaml:"port"`
	} `yaml:"auth_service"`
	Server struct {
		Address string `yaml:"address"`
		Port    int    `yaml:"port"`
	} `yaml:"server"`
}

func NewConfig() *Config {
	bytes, err := os.ReadFile("config.yml")
	if err != nil {
		logger.GetLogger().Fatal("Couldn't have read config file")
		return nil
	}
	var config Config
	err = yaml.Unmarshal(bytes, &config)
	if err != nil {
		logger.GetLogger().Fatal("Couldn't have unmarshalled config file",
			zap.String("err", err.Error()))
		return nil
	}
	return &config
}
