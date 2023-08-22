package config

import (
	"github.com/e1esm/LyrVibe/music-service/pkg/logger"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	GRPC struct {
		Address string `yaml:"addr"`
	} `yaml:"grpc"`
	MusicStorage struct {
		Address        string `yaml:"addr"`
		Port           int    `yaml:"port"`
		MaxConnections int    `yaml:"max_conn"`
		Database       string `yaml:"database"`
		User           string `yaml:"database_user"`
		Password       string `yaml:"database_password"`
		Name           string `yaml:"database_name"`
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
