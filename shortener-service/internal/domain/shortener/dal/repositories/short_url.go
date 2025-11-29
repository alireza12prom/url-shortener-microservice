package repositories

import (
	"github.com/gocql/gocql"
	"github.com/shortener-service/internal/common/exceptions"
	"github.com/shortener-service/internal/domain/shortener/dal/mappers"
	"github.com/shortener-service/internal/domain/shortener/dal/models"
	"github.com/shortener-service/internal/domain/shortener/entities"
	"github.com/shortener-service/internal/infrastructure/scylladb"
)

type ShortURLRepository struct {
	session *gocql.Session
}

func NewShortURLRepository(connection *scylladb.ScyllaDB) *ShortURLRepository {
	return &ShortURLRepository{
		session: connection.GetSession(),
	}
}

func (r *ShortURLRepository) Save(entity *entities.ShortURL) error {
	query := "INSERT INTO short_url (id, user_id, hash, endpoint, is_active, created_at, last_accessed_at) VALUES (?, ?, ?, ?, ?, ?, ?)"

	model := mappers.MapToShortURLModel(entity)

	err := r.session.Query(
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

func (r *ShortURLRepository) GetByHash(hash string) (*entities.ShortURL, error) {
	query := `
		SELECT 
			id,
			user_id, 
			endpoint, 
			hash, 
			is_active, 
			created_at, 
			last_accessed_at 
		FROM 
			short_url_by_hash 
		WHERE hash = ?`

	var result models.ShortURLModel
	err := r.session.Query(query, hash).Consistency(gocql.One).Scan(
		&result.ID,
		&result.UserID,
		&result.Endpoint,
		&result.Hash,
		&result.IsActive,
		&result.CreatedAt,
		&result.LastAccessedAt,
	)
	if err == gocql.ErrNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, exceptions.NewDatabaseException(err)
	}

	return mappers.MapToShortURLDomain(&result), nil
}
