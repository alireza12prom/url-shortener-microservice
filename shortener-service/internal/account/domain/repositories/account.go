package repositories

import "github.com/shortener-service/internal/account/domain/entities"

type AccountRepository interface {
	GetByID(id string) (*entities.AccountEntity, error)
	GetByEmail(email string) (*entities.AccountEntity, error)
	Save(user *entities.AccountEntity) error
}
