package service

import (
	"context"
	"github.com/e1esm/LyrVibe/artist-service/api/v1/proto"
)

type ArtistServiceProvider interface {
	New(*proto.VerificationRequest) (*proto.VerificationResponse, error)
}

type ArtistService struct {
	client proto.ArtistServiceClient
}

func NewArtistService(client proto.ArtistServiceClient) ArtistServiceProvider {
	return &ArtistService{client: client}
}

func (as *ArtistService) New(req *proto.VerificationRequest) (*proto.VerificationResponse, error) {
	resp, err := as.client.Verify(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
