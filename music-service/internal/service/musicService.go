package service

import (
	"context"
	"github.com/e1esm/LyrVibe/music-service/api/v1/proto"
	"github.com/e1esm/LyrVibe/music-service/internal/entity"
	"github.com/e1esm/LyrVibe/music-service/internal/repository"
)

type MusicServiceProvider interface {
	AddNew(context.Context, *proto.NewTrackRequest) (entity.TrackEntity, error)
}

type MusicService struct {
	Repository repository.Repository
}

func NewMusicService(repo repository.Repository) MusicServiceProvider {
	return &MusicService{Repository: repo}
}

func (ms *MusicService) AddNew(ctx context.Context, request *proto.NewTrackRequest) (entity.TrackEntity, error) {
	track := entity.NewTrackEntity(request)
	return ms.Repository.NewTrack(ctx, *track)
}
