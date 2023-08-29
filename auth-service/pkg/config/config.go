package config

import (
	"github.com/e1esm/LyrVibe/auth-service/pkg/logger"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	SessionsStorage struct {
		ContainerName string `yaml:"container_name"`
		Port          int    `yaml:"port"`
	} `yaml:"sessions_storage"`
	UsersStorage struct {
		ContainerName     string `yaml:"container_name"`
		Port              int    `yaml:"port,omitempty"`
		MaxConnectionPool int    `yaml:"max_connections"`
	} `yaml:"users_storage"`
	GRPC struct {
		Address string `json:"address"`
	} `json:"grpc"`
	HTTP struct {
		Address string `json:"address"`
	} `json:"http"`
}

func NewConfig() *Config {
	ymlContent, err := os.ReadFile("config.yml")
	if err != nil {
		logger.GetLogger().Error(err.Error())
		return nil
	}
	config := &Config{}
	err = yaml.Unmarshal(ymlContent, config)
	if err != nil {
		logger.GetLogger().Error(err.Error())
		return nil
	}
	return config
}
