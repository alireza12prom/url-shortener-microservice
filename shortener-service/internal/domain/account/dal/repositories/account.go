package repositories

import (
	"github.com/gocql/gocql"
	exceptions "github.com/shortener-service/internal/common/exceptions"
	"github.com/shortener-service/internal/domain/account/dal/mappers"
	"github.com/shortener-service/internal/domain/account/dal/models"
	"github.com/shortener-service/internal/domain/account/domain/entities"
	"github.com/shortener-service/internal/infrastructure/scylladb"
)

type AccountRepository struct {
	session *gocql.Session
}

func NewAccountRepository(connection *scylladb.ScyllaDB) *AccountRepository {
	return &AccountRepository{
		session: connection.GetSession(),
	}
}

func (r *AccountRepository) Save(account *entities.AccountEntity) error {
	query := "INSERT INTO account (id, name, username, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)"

	model := mappers.MapToAccountModel(account)

	err := r.session.Query(query, model.ID, model.Name, model.Username, model.Password, model.CreatedAt, model.UpdatedAt).Exec()
	if err != nil {
		return exceptions.NewDatabaseException(err)
	}

	return nil
}

func (r *AccountRepository) GetByUserId(userId string) (
	*entities.AccountEntity,
	error,
) {
	query := "SELECT id, name, username, password, created_at, updated_at FROM account WHERE id = ?"

	var result models.AccountModel
	err := r.session.Query(query, userId).Consistency(gocql.One).Scan(
		&result.ID,
		&result.Name,
		&result.Username,
		&result.Password,
		&result.CreatedAt,
		&result.UpdatedAt,
	)
	if err != nil {
		return nil, exceptions.NewDatabaseException(err)
	}

	return mappers.MapToAccountDomain(&result), nil
}

func (r *AccountRepository) GetByUsername(username string) (
	*entities.AccountEntity,
	error,
) {
	query := "SELECT id, name, username, password, created_at, updated_at FROM account_by_username WHERE username = ?"

	var result models.AccountModel
	err := r.session.Query(query, username).Consistency(gocql.One).Scan(
		&result.ID,
		&result.Name,
		&result.Username,
		&result.Password,
		&result.CreatedAt,
		&result.UpdatedAt,
	)
	if err == gocql.ErrNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, exceptions.NewDatabaseException(err)
	}

	return mappers.MapToAccountDomain(&result), nil
}

func (r *AccountRepository) IsUsernameUnique(username string) (
	bool,
	error,
) {
	query := "SELECT COUNT(*) FROM account_by_username WHERE username = ? LIMIT 1"

	var result int

	err := r.session.Query(query, username).Consistency(gocql.One).Scan(&result)
	if err != nil {
		return false, exceptions.NewDatabaseException(err)
	}

	return result == 0, nil
}
