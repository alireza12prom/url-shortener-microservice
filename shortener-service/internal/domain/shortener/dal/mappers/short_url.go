package mappers

import (
	"github.com/shortener-service/internal/domain/shortener/dal/models"
	"github.com/shortener-service/internal/domain/shortener/entities"
	"github.com/shortener-service/internal/domain/shortener/value-objects"
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
	ID, _ := valueobjects.NewID(m.ID)
	UserID, _ := valueobjects.NewID(m.UserID)
	Hash, _ := valueobjects.NewHash(m.Hash)
	Endpoint, _ := valueobjects.NewEndpoint(m.Endpoint)

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
