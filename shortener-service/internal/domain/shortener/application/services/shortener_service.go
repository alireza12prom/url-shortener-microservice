package services

import (
	"log"

	"github.com/google/uuid"
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
	entity, err := entities.NewShortURL(
		uuid.New().String(),
		command.UserID,
		command.Hash,
		command.Endpoint,
	)
	if err != nil {
		return nil, err
	}

	// TODO: check hash is unique

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
	entity, err := entities.NewShortURL(
		uuid.New().String(),
		command.UserID,
		s.hashReservationService.Reserve(command.Length),
		command.Endpoint,
	)
	if err != nil {
		return nil, err
	}

	// TODO: check hash is unique

	err = s.shortURLRepo.Save(&entity)
	if err != nil {
		return nil, err
	}

	err = s.eventPublisher.Publish(events.NewShortURLCreatedEvent(&entity))
	if err != nil {
		log.Println("Error publishing event:", err)
	}

	return &commands.ShortenRandomCommandOutput{
		ShortenURL: entity.Hash.GetValue(),
	}, nil
}
