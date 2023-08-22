package service

import "github.com/e1esm/LyrVibe/music-service/internal/repository"

type MusicServiceProvider interface {
}

type MusicService struct {
	Repository repository.Repository
}

func NewMusicService(repo repository.Repository) MusicServiceProvider {
	return &MusicService{Repository: repo}
}
