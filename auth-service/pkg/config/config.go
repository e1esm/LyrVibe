package config

type Config struct {
	SessionsStorage struct {
		ContainerName string `yaml:"container_name"`
		Port          int    `yaml:"port"`
	} `yaml:"sessions_storage"`
}
