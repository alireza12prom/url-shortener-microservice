package entities

import ValueObjects "github.com/warehouse-service/internal/domain/domain-events/domain/value-objects"

type DomainEventEntity struct {
	ID            *ValueObjects.ID
	Name          string
	Context       string
	Payload       any
	DateTime      string
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

	return &DomainEventEntity{
		ID:            ID,
		Name:          name,
		Context:       context,
		Payload:       payload,
		DateTime:      datetime,
		CorrelationID: CorrelationID,
	}, nil
}
