package repositories

import (
	"github.com/shortener-service/internal/account/dal/mappers"
	"github.com/shortener-service/internal/account/domain/entities"
	"github.com/shortener-service/internal/infrastructure/scylladb"
)

type AccountRepository struct {
	connection *scylladb.ScyllaDB
}

func NewAccountRepository(connection *scylladb.ScyllaDB) *AccountRepository {
	return &AccountRepository{
		connection: connection,
	}
}

func (r *AccountRepository) Save(account *entities.AccountEntity) error {
	query := "INSERT INTO accounts (id, username, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)"

	model := mappers.MapToAccountModel(account)

	err := r.connection.Session.Query(query, model.ID, model.Username, model.Email, model.Password, model.CreatedAt, model.UpdatedAt).Exec()
	if err != nil {
		return err
	}

	return nil
}
