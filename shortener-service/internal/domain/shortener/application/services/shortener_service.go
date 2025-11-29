package services

import (
	"log"

	"github.com/google/uuid"
	exceptions "github.com/shortener-service/internal/common/exceptions"
	"github.com/shortener-service/internal/domain/shortener/application/commands"
	"github.com/shortener-service/internal/domain/shortener/entities"
	"github.com/shortener-service/internal/domain/shortener/events"
	"github.com/shortener-service/internal/domain/shortener/interfaces"
	domain_services "github.com/shortener-service/internal/domain/shortener/services"
	"github.com/shortener-service/internal/lib"
)

type ShortenerService struct {
	shortURLRepo           interfaces.ShortURLRepository
	hashReservationService interfaces.HashReservationService
	eventPublisher         lib.EventPublisher
}

func NewShortenerService(
	shortURLRepo interfaces.ShortURLRepository,
	eventPublisher lib.EventPublisher,
) *ShortenerService {
	return &ShortenerService{
		shortURLRepo:           shortURLRepo,
		hashReservationService: domain_services.NewHashReservationService(),
		eventPublisher:         eventPublisher,
	}
}

func (s *ShortenerService) ShortenCustom(command *commands.ShortenCustomCommand) (
	*commands.ShortenCustomCommandOutput,
	error,
) {
	existing, err := s.shortURLRepo.GetByHash(command.Hash)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, exceptions.NewBusinessException(exceptions.HashAlreadyExists, map[string]any{
			"hash": command.Hash,
		})
	}

	entity, err := entities.NewShortURL(
		uuid.New().String(),
		command.UserID,
		command.Hash,
		command.Endpoint,
	)
	if err != nil {
		return nil, err
	}

	err = s.shortURLRepo.Save(&entity)
	if err != nil {
		return nil, err
	}

	err = s.eventPublisher.Publish(events.NewShortURLCreatedEvent(&entity))
	if err != nil {
		log.Println("Error publishing event:", err)
	}

	return &commands.ShortenCustomCommandOutput{
		ShortenURL: entity.Hash.GetValue(),
	}, nil
}

func (s *ShortenerService) ShortenRandom(command *commands.ShortenRandomCommand) (
	*commands.ShortenRandomCommandOutput,
	error,
) {
	const maxRetries = 10
	var entity entities.ShortURL
	var err error
	var success bool

	for attempt := 0; attempt < maxRetries; attempt++ {
		hash := s.hashReservationService.Reserve(command.Length)

		existing, err := s.shortURLRepo.GetByHash(hash)
		if err != nil {
			return nil, err
		}
		if existing != nil {
			continue
		}

		entity, err = entities.NewShortURL(
			uuid.New().String(),
			command.UserID,
			hash,
			command.Endpoint,
		)
		if err != nil {
			return nil, err
		}

		err = s.shortURLRepo.Save(&entity)
		if err != nil {
			return nil, err
		}

		success = true
		break
	}

	if !success {
		return nil, exceptions.NewBusinessException(exceptions.HashAlreadyExists, map[string]any{
			"message": "Failed to generate unique hash after multiple attempts",
		})
	}

	err = s.eventPublisher.Publish(events.NewShortURLCreatedEvent(&entity))
	if err != nil {
		log.Println("Error publishing event:", err)
	}

	return &commands.ShortenRandomCommandOutput{
		ShortenURL: entity.Hash.GetValue(),
	}, nil
}
