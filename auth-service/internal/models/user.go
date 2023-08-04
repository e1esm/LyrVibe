package models

import (
	"github.com/e1esm/LyrVibe/auth-service/api/v1/proto"
	"github.com/e1esm/LyrVibe/auth-service/pkg/hash"
	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID `json:"id,omitempty"`
	Username       string    `json:"username"`
	Password       string    `json:"password"`
	Role           Role      `json:"role"`
	Country        string    `json:"country"`
	FirstName      string    `json:"first_name"`
	SecondName     string    `json:"second_name"`
	ProfilePicture string    `json:"profile_picture,omitempty"`
}

func NewUser(pr *proto.SignUpRequest) *User {
	password, err := hash.GenerateHash(pr.Password)
	if err != nil {
		return nil
	}
	var userRole Role
	switch {
	case Role(pr.Role) == Admin:
		userRole = Admin
	case Role(pr.Role) == Artist:
		userRole = Artist
	default:
		userRole = Guest
	}
	return &User{
		ID:             uuid.New(),
		Username:       pr.Username,
		Password:       password,
		Role:           userRole,
		Country:        pr.Country,
		ProfilePicture: pr.Image,
		FirstName:      pr.FirstName,
		SecondName:     pr.SecondName,
	}
}
