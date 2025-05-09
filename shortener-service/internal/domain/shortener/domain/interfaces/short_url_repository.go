package interfaces

import "github.com/shortener-service/internal/domain/shortener/domain/entities"

type ShortURLRepository interface {
	Save(entity *entities.ShortURL) error
}
