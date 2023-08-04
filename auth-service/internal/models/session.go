package models

import "time"

type Session struct {
	StartTime time.Time    `json:"start_time"`
	Tokens    CachedTokens `json:"tokens"`
}
