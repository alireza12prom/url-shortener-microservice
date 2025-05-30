package services

import (
	"github.com/shortener-service/internal/common/exceptions"
	"github.com/shortener-service/internal/domain/redirector/application/commands"
	"github.com/shortener-service/internal/domain/redirector/domain/interfaces"
)

type RedirectorService struct {
	shortURLRepo interfaces.ShortURLRepository
}

func NewRedirectorService(shortURLRepo interfaces.ShortURLRepository) *RedirectorService {
	return &RedirectorService{
		shortURLRepo: shortURLRepo,
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

	return &commands.RedirectCommandOutput{
		Endpoint: shortURL.Endpoint.GetValue(),
	}, nil
}
