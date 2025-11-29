package events

import (
	"time"

	"github.com/google/uuid"
	"github.com/shortener-service/internal/domain/account/entities"
	"github.com/shortener-service/internal/lib"
)

type UsernameChangedPayload struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

func NewUsernameChangedEvent(entity *entities.AccountEntity) *lib.Event {
	return &lib.Event{
		ID:      entity.ID.GetValue(),
		Name:    "username_changed",
		Context: "account",
		Payload: PasswordChangedPayload{
			ID:       entity.ID.GetValue(),
			Username: entity.Username.GetValue(),
		},
		DateTime:      time.Now().Format(time.RFC3339),
		CorrelationID: uuid.New().String(), // FIXME: use request-id
	}
}
