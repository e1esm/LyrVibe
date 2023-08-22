package config

import (
	"github.com/e1esm/LyrVibe/music-service/pkg/logger"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	GRPC struct {
		address string `yaml:"addr"`
	} `yaml:"grpc"`
	MusicStorage struct {
		address        string `yaml:"addr"`
		port           int    `yaml:"port"`
		MaxConnections int    `yaml:"max_conn"`
	} `yaml:"music_storage"`
}

func NewConfig() *Config {
	content, err := os.ReadFile("config.yml")
	if err != nil {
		logger.GetLogger().Error(err.Error())
		return nil
	}
	cfg := &Config{}
	err = yaml.Unmarshal(content, cfg)
	if err != nil {
		logger.GetLogger().Error(err.Error())
		return nil
	}
	return cfg
}
