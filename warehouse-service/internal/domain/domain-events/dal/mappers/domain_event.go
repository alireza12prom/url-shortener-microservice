package mappers

import (
	"encoding/json"

	"github.com/warehouse-service/internal/domain/domain-events/dal/models"
	"github.com/warehouse-service/internal/domain/domain-events/domain/entities"
	ValueObjects "github.com/warehouse-service/internal/domain/domain-events/domain/value-objects"
)

func MapToAccountModel(e *entities.DomainEventEntity) *models.DomainEventModel {
	jsonPayload, _ := json.Marshal(e.Payload)

	return &models.DomainEventModel{
		ID:            e.ID.GetValue(),
		Name:          e.Name,
		Context:       e.Context,
		Payload:       string(jsonPayload),
		DateTime:      e.DateTime,
		CorrelationID: e.CorrelationID.GetValue(),
	}
}

func MapToAccountDomain(m *models.DomainEventModel) *entities.DomainEventEntity {
	ID, _ := ValueObjects.NewID(m.ID)
	CorrelationID, _ := ValueObjects.NewID(m.CorrelationID)

	return &entities.DomainEventEntity{
		ID:            ID,
		Name:          m.Name,
		Context:       m.Context,
		Payload:       m.Payload,
		DateTime:      m.DateTime,
		CorrelationID: CorrelationID,
	}
}
