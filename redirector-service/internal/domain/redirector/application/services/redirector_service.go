package services

import (
	"github.com/shortener-service/internal/common/exceptions"
	"github.com/shortener-service/internal/common/logger"
	"github.com/shortener-service/internal/domain/redirector/application/commands"
	"github.com/shortener-service/internal/domain/redirector/domain/events"
	"github.com/shortener-service/internal/domain/redirector/domain/interfaces"
	"github.com/shortener-service/internal/lib"
)

type RedirectorService struct {
	shortURLRepo   interfaces.ShortURLRepository
	eventPublisher lib.EventPublisher
	logger         *logger.Logger
}

func NewRedirectorService(
	shortURLRepo interfaces.ShortURLRepository,
	eventPublisher lib.EventPublisher,
) *RedirectorService {
	return &RedirectorService{
		shortURLRepo:   shortURLRepo,
		eventPublisher: eventPublisher,
		logger:         logger.NewLogger("redirector.service-domain"),
	}
}

func (Self *RedirectorService) Redirect(command *commands.RedirectCommand) (
	*commands.RedirectCommandOutput,
	error,
) {
	shortURL, err := Self.shortURLRepo.GetByHash(command.Hash)
	if err != nil {
		return nil, err
	}

	if shortURL == nil {
		return nil, exceptions.NewBusinessException(exceptions.ShortURLNotFound, nil)
	}

	err = Self.eventPublisher.Publish(events.NewRedirectHappenedEvent(shortURL))
	if err != nil {
		Self.logger.Error("Failed to publish domain event", logger.Fields{"Exception": err.Error()})
	}

	return &commands.RedirectCommandOutput{
		Endpoint: shortURL.Endpoint.GetValue(),
	}, nil
}
