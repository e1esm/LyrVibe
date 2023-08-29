package service

import (
	"context"
	"fmt"
	"github.com/e1esm/LyrVibe/artist-service/internal/models"
	"github.com/e1esm/LyrVibe/music-service/api/v1/proto"
)

type MusicServiceProvider interface {
	Release(context.Context, *models.Song) (*proto.NewTrackResponse, error)
}

type MusicService struct {
	musicClient proto.MusicServiceClient
}

func NewMusicService(client proto.MusicServiceClient) MusicServiceProvider {
	return &MusicService{musicClient: client}
}

func (m *MusicService) Release(ctx context.Context, song *models.Song) (*proto.NewTrackResponse, error) {
	releaseRequest := models.NewReleaseRequest(song)
	if releaseRequest == nil {
		return nil, fmt.Errorf("invalid Release request")
	}
	return m.musicClient.AddNewTrack(ctx, releaseRequest)
}
