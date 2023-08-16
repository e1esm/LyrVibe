package service

import (
	"context"
	"github.com/e1esm/LyrVibe/artist-service/api/v1/proto"
	"github.com/e1esm/LyrVibe/artist-service/internal/models"
	"github.com/e1esm/LyrVibe/artist-service/internal/repository"
)

type Service interface {
	AddArtist(context.Context, *proto.VerificationRequest) (*models.Artist, error)
}

type ArtistService struct {
	repository repository.Repository
}

func NewArtistService(repo repository.Repository) Service {
	return &ArtistService{repository: repo}
}

func (as *ArtistService) AddArtist(ctx context.Context, request *proto.VerificationRequest) (*models.Artist, error) {
	artist := models.NewArtist(request)
	err := as.repository.Add(ctx, artist)
	if err != nil {
		return nil, err
	}
	return artist, nil
}
