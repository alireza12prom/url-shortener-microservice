package interfaces

import "github.com/shortener-service/internal/domain/shortener/application/commands"

type ShortenerService interface {
	ShortenRandom(command *commands.ShortenRandomCommand) (*commands.ShortenRandomCommandOutput, error)
	ShortenCustom(command *commands.ShortenCustomCommand) (*commands.ShortenCustomCommandOutput, error)
}
