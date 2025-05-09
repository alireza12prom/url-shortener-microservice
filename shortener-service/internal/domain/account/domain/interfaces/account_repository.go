package interfaces

import "github.com/shortener-service/internal/domain/account/domain/entities"

type AccountRepository interface {
	Save(account *entities.AccountEntity) error
	GetByUserId(userId string) (*entities.AccountEntity, error)
	GetByUsername(userId string) (*entities.AccountEntity, error)
	IsUsernameUnique(username string) (bool, error)
}
