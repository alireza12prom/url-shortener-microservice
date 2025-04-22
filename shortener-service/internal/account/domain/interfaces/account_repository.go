package interfaces

import "github.com/shortener-service/internal/account/domain/entities"

type AccountRepository interface {
	Save(account *entities.AccountEntity) error
	GetByUserId(userId string) (*entities.AccountEntity, error)
	IsEmailUnique(email string) (bool, error)
	IsUsernameUnique(username string) (bool, error)
}
