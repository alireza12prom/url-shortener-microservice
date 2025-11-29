package interfaces

import "github.com/shortener-service/internal/domain/shortener/entities"

type ShortURLRepository interface {
	Save(entity *entities.ShortURL) error
	GetByHash(hash string) (*entities.ShortURL, error)
}
