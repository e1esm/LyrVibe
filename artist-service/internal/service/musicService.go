package service

import (
	"context"
	"fmt"
	artist "github.com/e1esm/LyrVibe/artist-service/api/v1/proto"
	"github.com/e1esm/LyrVibe/artist-service/internal/models"
	"github.com/e1esm/LyrVibe/music-service/api/v1/proto"
)

type MusicServiceProvider interface {
	Release(context.Context, *models.Song) (*proto.NewTrackResponse, error)
	DeleteTrack(context.Context, *artist.DeleteTrackRequest) (*artist.DeleteTrackResponse, error)
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

func (m *MusicService) DeleteTrack(ctx context.Context, deleteRequest *artist.DeleteTrackRequest) (*artist.DeleteTrackResponse, error) {
	resp, err := m.musicClient.DeleteTrack(ctx, &proto.DeleteRequest{AuthorId: deleteRequest.AuthorId, Title: deleteRequest.TrackTitle})
	if err != nil {
		return nil, err
	}
	
	return &artist.DeleteTrackResponse{
		Title:         resp.Title,
		RequestStatus: artist.RequestStatus_OK,
	}, nil
}
