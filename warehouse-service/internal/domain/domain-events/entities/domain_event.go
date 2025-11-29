package entities

import (
	"time"

	ValueObjects "github.com/warehouse-service/internal/domain/domain-events/value-objects"
)

type DomainEventEntity struct {
	ID            *ValueObjects.ID
	Name          string
	Context       string
	Payload       any
	DateTime      *time.Time
	CorrelationID *ValueObjects.ID
}

func NewDomainEventEntity(
	id,
	name,
	context string,
	payload any,
	correlationId string,
	datetime string,
) (*DomainEventEntity, error) {
	ID, err := ValueObjects.NewID(id)
	if err != nil {
		return nil, err
	}

	CorrelationID, err := ValueObjects.NewID(correlationId)
	if err != nil {
		return nil, err
	}

	DateTime, error := time.Parse(time.RFC3339, datetime)
	if error != nil {
		return nil, error
	}

	return &DomainEventEntity{
		ID:            ID,
		Name:          name,
		Context:       context,
		Payload:       payload,
		DateTime:      &DateTime,
		CorrelationID: CorrelationID,
	}, nil
}
