package entities

import (
	"time"

	ValueObjects "github.com/shortener-service/internal/domain/redirector/value-objects"
)

type ShortURL struct {
	ID             *ValueObjects.ID
	UserID         *ValueObjects.ID
	Hash           *ValueObjects.Hash
	IsActive       bool
	Endpoint       *ValueObjects.Endpoint
	CreatedAt      time.Time
	LastAccessedAt time.Time
}

func NewShortURL(id, userId, hash, endpoint string) (ShortURL, error) {
	ID, err := ValueObjects.NewID(id)
	if err != nil {
		return ShortURL{}, err
	}

	UserID, err := ValueObjects.NewID(userId)
	if err != nil {
		return ShortURL{}, err
	}

	Hash, err := ValueObjects.NewHash(hash)
	if err != nil {
		return ShortURL{}, err
	}

	Endpoint, err := ValueObjects.NewEndpoint(endpoint)
	if err != nil {
		return ShortURL{}, err
	}

	return ShortURL{
		ID:             ID,
		UserID:         UserID,
		Hash:           Hash,
		IsActive:       true,
		Endpoint:       Endpoint,
		CreatedAt:      time.Now(),
		LastAccessedAt: time.Now(),
	}, nil
}
