package mappers

import (
	"encoding/json"
	"time"

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
		DateTime:      e.DateTime.Format("2006-01-02 15:04:05"),
		CorrelationID: e.CorrelationID.GetValue(),
	}
}

func MapToAccountDomain(m *models.DomainEventModel) *entities.DomainEventEntity {
	ID, _ := ValueObjects.NewID(m.ID)
	CorrelationID, _ := ValueObjects.NewID(m.CorrelationID)
	DateTime, _ := time.Parse("2006-01-02 15:04:05", m.DateTime)

	return &entities.DomainEventEntity{
		ID:            ID,
		Name:          m.Name,
		Context:       m.Context,
		Payload:       m.Payload,
		DateTime:      &DateTime,
		CorrelationID: CorrelationID,
	}
}
