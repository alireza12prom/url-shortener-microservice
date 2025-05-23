package events

import (
	"time"

	"github.com/google/uuid"
	"github.com/shortener-service/internal/domain/account/domain/entities"
	"github.com/shortener-service/internal/lib"
)

type AccountCreatedPayload struct {
	AccountId string `json:"shortener_id"`
	Username  string `json:"username"`
}

func NewAccountCreatedEvent(entity *entities.AccountEntity) *lib.Event {
	return &lib.Event{
		ID:      entity.ID.GetValue(),
		Name:    "account_created",
		Context: "account",
		Payload: AccountCreatedPayload{
			AccountId: entity.ID.GetValue(),
			Username:  entity.Username.GetValue(),
		},
		DateTime:      time.Now().Format(time.RFC3339),
		CorrelationID: uuid.New().String(),
	}
}
