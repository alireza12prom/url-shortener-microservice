package repositories

import "github.com/shortener-service/internal/infrastructure/scylladb"

type ShortURLRepository struct {
	connection *scylladb.ScyllaDB
}

func NewAccountRepository(connection *scylladb.ScyllaDB) *ShortURLRepository {
	return &ShortURLRepository{
		connection: connection,
	}
}
