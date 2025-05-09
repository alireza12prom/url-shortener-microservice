package mappers

import (
	"github.com/shortener-service/internal/domain/shortener/dal/models"
	"github.com/shortener-service/internal/domain/shortener/domain/entities"
	ValueObjects "github.com/shortener-service/internal/domain/shortener/domain/value-objects"
)

func MapToShortURLModel(e *entities.ShortURL) *models.ShortURLModel {
	return &models.ShortURLModel{
		ID:             e.ID.GetValue(),
		UserID:         e.UserID.GetValue(),
		Hash:           e.Hash.GetValue(),
		Endpoint:       e.Endpoint.GetValue(),
		IsActive:       e.IsActive,
		CreatedAt:      e.CreatedAt,
		LastAccessedAt: e.LastAccessedAt,
	}
}

func MapToShortURLDomain(m *models.ShortURLModel) *entities.ShortURL {
	ID, _ := ValueObjects.NewID(m.ID)
	UserID, _ := ValueObjects.NewID(m.UserID)
	Hash, _ := ValueObjects.NewHash(m.Hash)
	Endpoint, _ := ValueObjects.NewEndpoint(m.Endpoint)

	return &entities.ShortURL{
		ID:             ID,
		UserID:         UserID,
		Hash:           Hash,
		Endpoint:       Endpoint,
		IsActive:       m.IsActive,
		CreatedAt:      m.CreatedAt,
		LastAccessedAt: m.LastAccessedAt,
	}
}
