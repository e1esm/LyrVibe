package service

import "github.com/e1esm/LyrVibe/artist-service/internal/repository"

type Service interface {
}

type ArtistService struct {
	repository repository.Repository
}

func NewArtistService(repo repository.Repository) Service {
	return ArtistService{repository: repo}
}
