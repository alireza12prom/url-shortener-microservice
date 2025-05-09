package services

import (
	"github.com/google/uuid"
	"github.com/shortener-service/internal/domain/shortener/application/commands"
	"github.com/shortener-service/internal/domain/shortener/domain/entities"
	"github.com/shortener-service/internal/domain/shortener/domain/interfaces"
	domain_services "github.com/shortener-service/internal/domain/shortener/domain/services"
)

type ShortenerService struct {
	shortURLRepo           interfaces.ShortURLRepository
	hashReservationService interfaces.HashReservationService
}

func NewShortenerService(shortURLRepo interfaces.ShortURLRepository) *ShortenerService {
	return &ShortenerService{
		shortURLRepo:           shortURLRepo,
		hashReservationService: domain_services.NewHashReservationService(),
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

	return &commands.ShortenCustomCommandOutput{}, nil
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

	return &commands.ShortenRandomCommandOutput{}, nil
}
