package models

import "time"

type ShortURLModel struct {
	ID             string
	UserID         string
	Hash           string
	Endpoint       string
	IsActive       bool
	CreatedAt      time.Time
	LastAccessedAt time.Time
}
