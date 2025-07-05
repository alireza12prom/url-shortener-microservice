package handlers

import (
	"encoding/json"

	"github.com/warehouse-service/internal/common/logger"
	"github.com/warehouse-service/internal/domain/domain-events/application/commands"
	"github.com/warehouse-service/internal/lib"
)

type CaptureEventConsumer struct {
	Logger  *logger.Logger
	Service lib.ApplicationService
}

func NewCaptureEventConsumer(service lib.ApplicationService) *CaptureEventConsumer {
	return &CaptureEventConsumer{
		Service: service,
		Logger:  logger.NewLogger("CaptureEventConsumer"),
	}
}

func (Self *CaptureEventConsumer) Handle(input []byte) error {
	var command commands.CaptureEventCommand
	if err := json.Unmarshal(input, &command); err != nil {
		return err
	}

	Self.Logger.Debug("Command", logger.Fields{
		"command": command,
	})

	err := Self.Service.Exec(&command)
	if err != nil {
		return err
	}

	return nil
}
