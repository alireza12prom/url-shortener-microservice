package interfaces

import "github.com/shortener-service/internal/domain/redirector/entities"

type ShortURLRepository interface {
	GetByHash(hash string) (*entities.ShortURL, error)
}
