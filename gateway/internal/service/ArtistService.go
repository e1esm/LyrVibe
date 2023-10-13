package service

import (
	"context"
	"github.com/e1esm/LyrVibe/artist-service/api/v1/proto"
	"github.com/e1esm/LyrVibe/gateway/internal/registrator"
	"github.com/e1esm/LyrVibe/gateway/pkg/config"
)

type ArtistServiceProvider interface {
	New(*proto.VerificationRequest) (*proto.VerificationResponse, error)
	ReleaseTrack(*proto.NewTrackRequest) (*proto.NewTrackResponse, error)
	DeleteTrack(*proto.DeleteTrackRequest) (*proto.DeleteTrackResponse, error)
	ReleaseAlbum(*proto.NewAlbumRequest) (*proto.NewAlbumResponse, error)
}

type ArtistService struct {
	client proto.ArtistServiceClient
}

func NewArtistService(cfg config.Config) ArtistServiceProvider {
	return &ArtistService{client: registrator.RegisterArtistService(&cfg)}
}

func (as *ArtistService) New(req *proto.VerificationRequest) (*proto.VerificationResponse, error) {
	return as.client.Verify(context.Background(), req)
}

func (as *ArtistService) ReleaseTrack(req *proto.NewTrackRequest) (*proto.NewTrackResponse, error) {
	return as.client.AddTrack(context.Background(), req)
}

func (as *ArtistService) DeleteTrack(req *proto.DeleteTrackRequest) (*proto.DeleteTrackResponse, error) {
	return as.client.DeleteTrack(context.Background(), req)
}

func (as *ArtistService) ReleaseAlbum(req *proto.NewAlbumRequest) (*proto.NewAlbumResponse, error) {
	return as.client.AddAlbum(context.Background(), req)
}
