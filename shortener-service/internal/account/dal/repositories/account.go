package repositories

import (
	"github.com/gocql/gocql"
	"github.com/shortener-service/internal/account/dal/mappers"
	"github.com/shortener-service/internal/account/dal/models"
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
	query := "INSERT INTO account (id, username, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)"

	model := mappers.MapToAccountModel(account)

	err := r.connection.Session.Query(query, model.ID, model.Username, model.Email, model.Password, model.CreatedAt, model.UpdatedAt).Exec()
	if err != nil {
		return err
	}

	return nil
}

func (r *AccountRepository) GetByUserId(userId string) (*entities.AccountEntity, error) {
	query := "SELECT id, username, email, password, created_at, updated_at FROM account WHERE id = ?"

	var result models.AccountModel
	err := r.connection.Session.Query(query, userId).Consistency(gocql.One).Scan(
		&result.ID,
		&result.Username,
		&result.Email,
		&result.Password,
		&result.CreatedAt,
		&result.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return mappers.MapToAccountDomain(&result), nil
}

func (r *AccountRepository) IsEmailUnique(email string) (bool, error) {
	query := "SELECT COUNT(*) FROM account_by_email WHERE email = ? LIMIT 1"

	var result int

	err := r.connection.Session.Query(query, email).Consistency(gocql.One).Scan(&result)
	if err != nil {
		return false, err
	}

	return result == 0, nil
}

func (r *AccountRepository) IsUsernameUnique(email string) (bool, error) {
	query := "SELECT COUNT(*) FROM account_by_username WHERE username = ? LIMIT 1"

	var result int

	err := r.connection.Session.Query(query, email).Consistency(gocql.One).Scan(&result)
	if err != nil {
		return false, err
	}

	return result == 0, nil
}
