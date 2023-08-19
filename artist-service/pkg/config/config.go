package config

import (
	"github.com/e1esm/LyrVibe/artist-service/pkg/logger"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Server struct {
		Address string `yaml:"address"`
	} `yaml:"server"`
	ArtistStorage struct {
		ContainerName    string `yaml:"container_name"`
		Port             int    `yaml:"port"`
		DatabaseName     string `yaml:"database_name"`
		DatabaseUser     string `yaml:"database_user"`
		DatabasePassword string `yaml:"database_password"`
		Database         string `yaml:"database"`
		MaxConnections   int    `yaml:"max_connections"`
	} `yaml:"artist_storage"`
	MusicServiceServer struct {
		ContainerName string `yaml:"container_name"`
		Port          int    `yaml:"port"`
	} `yaml:"music_service"`
	AuthService struct {
		ContainerName string `yaml:"container_name"`
		Port          int    `yaml:"port"`
	} `yaml:"auth_service"`
}

func NewConfig() *Config {
	config := &Config{}
	content, err := os.ReadFile("config.yml")
	if err != nil {
		logger.Logger.Error(err.Error())
		return nil
	}
	err = yaml.Unmarshal(content, config)
	if err != nil {
		logger.Logger.Error(err.Error())
		return nil
	}
	return config
}
