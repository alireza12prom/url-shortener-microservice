package repositories

import (
	"github.com/gocql/gocql"
	"github.com/shortener-service/internal/common/exceptions"
	"github.com/shortener-service/internal/domain/redirector/dal/mappers"
	"github.com/shortener-service/internal/domain/redirector/dal/models"
	"github.com/shortener-service/internal/domain/redirector/entities"
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

func (r *ShortURLRepository) GetByHash(hash string) (
	*entities.ShortURL,
	error,
) {
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
