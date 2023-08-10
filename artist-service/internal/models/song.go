package models

import (
	"github.com/google/uuid"
	"time"
)

type Song struct {
	ID       uuid.UUID     `json:"id"`
	Title    string        `json:"title"`
	Genre    string        `json:"genre"`
	Cover    string        `json:"cover"`
	Lyrics   string        `json:"lyrics"`
	Duration time.Duration `json:"duration"`
	Feature  []Artist      `json:"feature"`
	Country  string        `json:"country"`
	YTLink   string        `json:"youtube_link,omitempty"`
	Views    int           `json:"views,omitempty"`
}
