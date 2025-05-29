package events

import (
	"time"

	"github.com/google/uuid"
	"github.com/shortener-service/internal/domain/shortener/domain/entities"
	"github.com/shortener-service/internal/lib"
)

type ShortURLCreatedPayload struct{}

func NewShortURLCreatedEvent(entity *entities.ShortURL) *lib.Event {
	return &lib.Event{
		ID:            entity.ID.GetValue(),
		Name:          "account_created",
		Context:       "account",
		Payload:       ShortURLCreatedPayload{},
		DateTime:      time.Now().Format(time.RFC3339),
		CorrelationID: uuid.New().String(),
	}
}
