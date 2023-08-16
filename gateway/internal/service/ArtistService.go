package service

import (
	"context"
	"github.com/e1esm/LyrVibe/artist-service/api/v1/proto"
	"github.com/e1esm/LyrVibe/gateway/internal/registrator"
	"github.com/e1esm/LyrVibe/gateway/pkg/config"
)

type ArtistServiceProvider interface {
	New(*proto.VerificationRequest) (*proto.VerificationResponse, error)
}

type ArtistService struct {
	client proto.ArtistServiceClient
}

func NewArtistService(cfg config.Config) ArtistServiceProvider {
	return &ArtistService{client: registrator.RegisterArtistService(&cfg)}
}

func (as *ArtistService) New(req *proto.VerificationRequest) (*proto.VerificationResponse, error) {
	resp, err := as.client.Verify(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
