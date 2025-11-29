package events

import (
	"time"

	"github.com/google/uuid"
	"github.com/shortener-service/internal/domain/redirector/entities"
	"github.com/shortener-service/internal/lib"
)

type RedirectHappenedPayload struct {
	ShortenerID string `json:"shortener_id"`
	UserID      string `json:"user_id"`
	Hash        string `json:"hash"`
	Endpoint    string `json:"endpoint"`
}

func NewRedirectHappenedEvent(entity *entities.ShortURL) *lib.Event {
	return &lib.Event{
		ID:      entity.ID.GetValue(),
		Name:    "redirect_happened",
		Context: "redirect",
		Payload: RedirectHappenedPayload{
			ShortenerID: entity.ID.GetValue(),
			UserID:      entity.UserID.GetValue(),
			Hash:        entity.Hash.GetValue(),
			Endpoint:    entity.Endpoint.GetValue(),
		},
		DateTime:      time.Now().Format(time.RFC3339),
		CorrelationID: uuid.New().String(),
	}
}
