package entity

import (
	"github.com/e1esm/LyrVibe/music-service/api/v1/proto"
	"github.com/google/uuid"
	"strings"
	"time"
)

type TrackEntity struct {
	ID        uuid.UUID
	Data      *proto.NewTrackRequest
	CreatedAt time.Time
	AlbumID   uuid.UUID
}

func NewTrackEntity(track *proto.NewTrackRequest) *TrackEntity {
	track.ReleaseDate = track.ReleaseDate[:strings.Index(track.ReleaseDate, " ")]
	return &TrackEntity{
		uuid.New(),
		track,
		time.Now().UTC().Round(time.Microsecond),
		uuid.UUID{},
	}
}
