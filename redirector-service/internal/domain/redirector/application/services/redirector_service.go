package services

import (
	"github.com/shortener-service/internal/domain/redirector/application/commands"
	"github.com/shortener-service/internal/domain/redirector/domain/interfaces"
)

type RedirectorService struct {
	shortURLRepo interfaces.ShortURLRepository
}

func (Self *RedirectorService) Redirect(command commands.RedirectCommand) (
	*commands.RedirectCommandOutput,
	error,
) {
	return nil, nil
}
