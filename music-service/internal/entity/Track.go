package entity

import (
	"github.com/e1esm/LyrVibe/music-service/api/v1/proto"
	"github.com/google/uuid"
	"time"
)

type TrackEntity struct {
	ID        uuid.UUID
	Data      *proto.NewTrackRequest
	CreatedAt time.Time
}

func NewTrackEntity(track *proto.NewTrackRequest) *TrackEntity {
	return &TrackEntity{
		uuid.New(),
		track,
		time.Now(),
	}
}
