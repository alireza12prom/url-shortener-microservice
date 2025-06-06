package events

import (
	"time"

	"github.com/google/uuid"
	"github.com/shortener-service/internal/domain/account/domain/entities"
	"github.com/shortener-service/internal/lib"
)

type LoginHappenedPayload struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

func NewLoginHappenedEvent(entity *entities.AccountEntity) *lib.Event {
	return &lib.Event{
		ID:      entity.ID.GetValue(),
		Name:    "login_happened",
		Context: "account",
		Payload: LoginHappenedPayload{
			ID:       entity.ID.GetValue(),
			Username: entity.Username.GetValue(),
		},
		DateTime:      time.Now().Format(time.RFC3339),
		CorrelationID: uuid.New().String(),
	}
}
