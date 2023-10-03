package entity

import (
	"github.com/e1esm/LyrVibe/music-service/api/v1/proto"
	"github.com/google/uuid"
)

type AlbumEntity struct {
	ID     uuid.UUID
	Title  string
	Tracks []TrackEntity
}

func NewAlbum(request *proto.NewAlbumRequest) *AlbumEntity {
	return &AlbumEntity{
		ID:     uuid.New(),
		Title:  request.Title,
		Tracks: getTracks(request.Tracks),
	}
}
