package models

type Role int

const (
	Guest Role = iota
	Artist
	Admin
)
