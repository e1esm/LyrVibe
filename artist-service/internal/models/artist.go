package models

import (
	"github.com/e1esm/LyrVibe/artist-service/api/v1/proto"
	"github.com/e1esm/LyrVibe/artist-service/pkg/logger"
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
	id, err := uuid.Parse(request.Id)
	if err != nil {
		logger.Logger.Error(err.Error())
		return nil
	}
	return &Artist{
		ID:         id,
		Username:   request.Username,
		Country:    request.Country,
		FirstName:  request.FirstName,
		SecondName: request.SecondName,
	}
}
