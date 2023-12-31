package models

import (
	"fmt"
	"github.com/e1esm/LyrVibe/auth-service/api/v1/proto"
	"github.com/e1esm/LyrVibe/auth-service/pkg/hash"
	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID `json:"id,omitempty"`
	Username       string    `json:"username"`
	Password       string    `json:"password"`
	Role           Role      `json:"role"`
	ProfilePicture []byte    `json:"profile_picture,omitempty"`
}

func NewUser(pr *proto.SignUpRequest) *User {
	password := fmt.Sprintf("%x", hash.GenerateHash(pr.Password))
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
		ProfilePicture: pr.Image,
	}
}
