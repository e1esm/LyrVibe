package entity

import (
	"github.com/e1esm/LyrVibe/music-service/api/v1/proto"
	"time"
)

type TrackEntity struct {
	*proto.NewTrackRequest
	CreatedAt time.Time
}

func NewTrackEntity(track *proto.NewTrackRequest) *TrackEntity {
	return &TrackEntity{
		track,
		time.Now(),
	}
}
