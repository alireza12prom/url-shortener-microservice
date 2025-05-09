package repositories

import (
	"github.com/shortener-service/internal/common/exceptions"
	"github.com/shortener-service/internal/domain/shortener/dal/mappers"
	"github.com/shortener-service/internal/domain/shortener/domain/entities"
	"github.com/shortener-service/internal/infrastructure/scylladb"
)

type ShortURLRepository struct {
	connection *scylladb.ScyllaDB
}

func NewShortURLRepository(connection *scylladb.ScyllaDB) *ShortURLRepository {
	return &ShortURLRepository{
		connection: connection,
	}
}

func (s *ShortURLRepository) Save(entity *entities.ShortURL) error {
	query := "INSERT INTO account (id, user_id, hash, endpoint, is_active, created_at, last_accessed_at) VALUES (?, ?, ?, ?, ?, ?, ?)"

	model := mappers.MapToShortURLModel(entity)

	err := s.connection.Session.Query(
		query,
		model.ID,
		model.UserID,
		model.Hash,
		model.Endpoint,
		model.IsActive,
		model.CreatedAt,
		model.LastAccessedAt,
	).Exec()
	if err != nil {
		return exceptions.NewDatabaseException(err)
	}

	return nil
}
