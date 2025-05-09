package services

import (
	"github.com/shortener-service/internal/domain/shortener/application/commands"
	"github.com/shortener-service/internal/domain/shortener/domain/interfaces"
)

type ShortenerService struct {
	shortURLRepo interfaces.ShortURLRepository
}

func NewShortenerService(shortURLRepo interfaces.ShortURLRepository) *ShortenerService {
	return &ShortenerService{
		shortURLRepo: shortURLRepo,
	}
}

func (s *ShortenerService) ShortenCustom(command *commands.ShortenCustomCommand) (
	*commands.ShortenCustomCommandOutput,
	error,
) {
	return &commands.ShortenCustomCommandOutput{}, nil
}

func (s *ShortenerService) ShortenRandom(command *commands.ShortenRandomCommand) (
	*commands.ShortenRandomCommandOutput,
	error,
) {
	return &commands.ShortenRandomCommandOutput{}, nil
}
