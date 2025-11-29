package services

import (
	"github.com/warehouse-service/internal/common/exceptions"
	"github.com/warehouse-service/internal/domain/domain-events/application/commands"
	"github.com/warehouse-service/internal/domain/domain-events/entities"
	"github.com/warehouse-service/internal/domain/domain-events/interfaces"
)

type CaptureEventService struct {
	EventRepo interfaces.DomainEventRepository
}

func (Self *CaptureEventService) Exec(command interface{}) error {
	cmd, ok := command.(*commands.CaptureEventCommand)
	if !ok {
		return exceptions.NewBusinessException(exceptions.InvalidCommandType, nil)
	}

	entity, err := entities.NewDomainEventEntity(
		cmd.ID,
		cmd.Name,
		cmd.Context,
		cmd.Payload,
		cmd.CorrelationID,
		cmd.DateTime,
	)
	if err != nil {
		return err
	}

	err = Self.EventRepo.Save(entity)
	if err != nil {
		return err
	}

	return nil
}
