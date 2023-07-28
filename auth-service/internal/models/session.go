package models

import "github.com/google/uuid"

type Session struct {
	SessionID uuid.UUID `json:"session_id"`
	UserID    uuid.UUID `json:"user_id"`
	UserRole  Role      `json:"user_role"`
}
