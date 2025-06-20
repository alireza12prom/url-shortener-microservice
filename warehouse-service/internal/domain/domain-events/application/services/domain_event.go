package services

import (
	"github.com/warehouse-service/internal/domain/domain-events/application/commands"
	"github.com/warehouse-service/internal/domain/domain-events/domain/entities"
	"github.com/warehouse-service/internal/domain/domain-events/domain/interfaces"
)

type DomainEventService struct {
	eventRepo interfaces.DomainEventRepository
}

func (Self *DomainEventService) CaptureEvent(
	command *commands.CaptureEventCommand,
) error {
	entity, err := entities.NewDomainEventEntity(
		command.ID,
		command.Name,
		command.Context,
		command.Payload,
		command.CorrelationID,
		command.DateTime,
	)
	if err != nil {
		return err
	}

	err = Self.eventRepo.Save(entity)
	if err != nil {
		return err
	}

	return nil
}
