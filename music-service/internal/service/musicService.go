package service

import (
	"context"
	"fmt"
	"github.com/e1esm/LyrVibe/music-service/api/v1/proto"
	"github.com/e1esm/LyrVibe/music-service/internal/entity"
	"github.com/e1esm/LyrVibe/music-service/internal/repository"
)

type MusicServiceProvider interface {
	AddNewTrack(context.Context, *proto.NewTrackRequest) (entity.TrackEntity, error)
	Delete(context.Context, *proto.DeleteRequest) (*proto.DeleteResponse, error)
	AddNewAlbum(context.Context, *proto.NewAlbumRequest) (entity.AlbumEntity, error)
}

type MusicService struct {
	Repository repository.Repository
}

func NewMusicService(repo repository.Repository) MusicServiceProvider {
	return &MusicService{Repository: repo}
}

func (ms *MusicService) AddNewTrack(ctx context.Context, request *proto.NewTrackRequest) (entity.TrackEntity, error) {
	track := entity.NewTrackEntity(request)
	return ms.Repository.NewTrack(ctx, *track)
}

func (ms *MusicService) Delete(ctx context.Context, request *proto.DeleteRequest) (*proto.DeleteResponse, error) {
	resp, err := ms.Repository.DeleteTrack(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("track wasn't deleted: %v", err)
	}
	return resp, nil
}

func (ms *MusicService) AddNewAlbum(ctx context.Context, request *proto.NewAlbumRequest) (entity.AlbumEntity, error) {
	return ms.Repository.NewAlbum(ctx, *entity.NewAlbum(request))
}
