package events

import (
	"time"

	"github.com/google/uuid"
	"github.com/shortener-service/internal/domain/account/domain/entities"
	"github.com/shortener-service/internal/lib"
)

type PasswordChangedPayload struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

func NewPasswordChangedEvent(entity *entities.AccountEntity) *lib.Event {
	return &lib.Event{
		ID:      entity.ID.GetValue(),
		Name:    "password_changed",
		Context: "account",
		Payload: PasswordChangedPayload{
			ID:       entity.ID.GetValue(),
			Username: entity.ID.GetValue(),
		},
		DateTime:      time.Now().Format(time.RFC3339),
		CorrelationID: uuid.New().String(),
	}
}
