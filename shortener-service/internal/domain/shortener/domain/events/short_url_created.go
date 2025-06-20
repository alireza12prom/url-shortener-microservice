package events

import (
	"time"

	"github.com/google/uuid"
	"github.com/shortener-service/internal/domain/shortener/domain/entities"
	"github.com/shortener-service/internal/lib"
)

type ShortURLCreatedPayload struct {
	ID       string `json:"id"`
	UserID   string `json:"user_id"`
	Hash     string `json:"hash"`
	Endpoint string `json:"endpoint"`
}

func NewShortURLCreatedEvent(entity *entities.ShortURL) *lib.Event {
	return &lib.Event{
		ID:      entity.ID.GetValue(),
		Name:    "short_url_created",
		Context: "shortener",
		Payload: ShortURLCreatedPayload{
			ID:       entity.ID.GetValue(),
			UserID:   entity.UserID.GetValue(),
			Hash:     entity.Hash.GetValue(),
			Endpoint: entity.Endpoint.GetValue(),
		},
		DateTime:      time.Now().Format(time.RFC3339),
		CorrelationID: uuid.New().String(),
	}
}
