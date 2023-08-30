package models

import (
	artist "github.com/e1esm/LyrVibe/artist-service/api/v1/proto"
	"github.com/e1esm/LyrVibe/artist-service/pkg/logger"
	music "github.com/e1esm/LyrVibe/music-service/api/v1/proto"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"time"
)

type Song struct {
	Title       string        `json:"title"`
	Genre       string        `json:"genre"`
	Cover       []byte        `json:"cover"`
	Lyrics      []string      `json:"lyrics"`
	Duration    time.Duration `json:"duration"`
	Feature     []string      `json:"feature"`
	Country     string        `json:"country"`
	VideoLink   string        `json:"video_link,omitempty"`
	ReleaseDate time.Time     `json:"release_date,omitempty"`
	ArtistID    uuid.UUID     `json:"artist_id"`
}

func NewReleaseRequest(song *Song) *music.NewTrackRequest {
	var songGenre int32 = -1
	for k, v := range music.Genre_name {
		if v == song.Genre {
			songGenre = k
		}
	}
	if songGenre == -1 {
		return nil
	}
	return &music.NewTrackRequest{
		Title:       song.Title,
		Genre:       music.Genre(songGenre),
		Cover:       song.Cover,
		Lyrics:      song.Lyrics,
		Duration:    song.Duration.String(),
		Feature:     song.Feature,
		Country:     song.Country,
		VideoLink:   song.VideoLink,
		ReleaseDate: song.ReleaseDate.String(),
	}
}

func NewSong(request *artist.NewTrackRequest) *Song {
	duration, err := time.ParseDuration(request.Duration)
	if err != nil {
		logger.GetLogger().Error("Can't convert string to duration",
			zap.String("err", err.Error()))
		return nil
	}
	release, err := time.Parse(time.UnixDate, request.ReleaseDate)
	if err != nil {
		logger.GetLogger().Error("Can't convert string to time",
			zap.String("err", err.Error()))
		return nil
	}

	artistID, err := uuid.Parse(request.ArtistId)
	if err != nil {
		logger.GetLogger().Error("Cant parse uuid",
			zap.String("err", err.Error()))
		return nil
	}
	return &Song{
		Title:       request.Title,
		Genre:       request.Genre.String(),
		Cover:       request.Cover,
		Lyrics:      request.Lyrics,
		Duration:    duration,
		Feature:     request.Feature,
		Country:     request.Country,
		VideoLink:   request.VideoLink,
		ReleaseDate: release,
		ArtistID:    artistID,
	}
}
