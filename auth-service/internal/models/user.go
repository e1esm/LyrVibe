package models

import "github.com/google/uuid"

type User struct {
	ID             uuid.UUID `json:"id,omitempty"`
	Username       string    `json:"username"`
	Password       string    `json:"password"`
	Role           Role      `json:"role"`
	Country        string    `json:"country"`
	FirstName      string    `json:"first_name"`
	SecondName     string    `json:"second_name"`
	ProfilePicture string    `json:"profile_picture"`
}
