package models

import (
	"github.com/e1esm/LyrVibe/artist-service/api/v1/proto"
	"github.com/google/uuid"
)

type Artist struct {
	ID         uuid.UUID `json:"id,omitempty"`
	Username   string    `json:"username"`
	Country    string    `json:"country" `
	FirstName  string    `json:"first_name"`
	SecondName string    `json:"second_name"`
	Views      uint      `json:"overall_views,omitempty"`
}

func NewArtist(request *proto.VerificationRequest) *Artist {
	return &Artist{
		ID:         uuid.New(),
		Username:   request.Username,
		Country:    request.Country,
		FirstName:  request.FirstName,
		SecondName: request.SecondName,
	}
}
