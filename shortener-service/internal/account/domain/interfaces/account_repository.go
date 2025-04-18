package interfaces

import "github.com/shortener-service/internal/account/domain/entities"

type AccountRepository interface {
	Save(account *entities.AccountEntity) error
}
