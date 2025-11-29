package entities

import (
	"time"

	"github.com/shortener-service/internal/domain/shortener/value-objects"
)

type ShortURL struct {
	ID             *valueobjects.ID
	UserID         *valueobjects.ID
	Hash           *valueobjects.Hash
	IsActive       bool
	Endpoint       *valueobjects.Endpoint
	CreatedAt      time.Time
	LastAccessedAt time.Time
}

func NewShortURL(id, userId, hash, endpoint string) (ShortURL, error) {
	ID, err := valueobjects.NewID(id)
	if err != nil {
		return ShortURL{}, err
	}

	UserID, err := valueobjects.NewID(userId)
	if err != nil {
		return ShortURL{}, err
	}

	Hash, err := valueobjects.NewHash(hash)
	if err != nil {
		return ShortURL{}, err
	}

	Endpoint, err := valueobjects.NewEndpoint(endpoint)
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
