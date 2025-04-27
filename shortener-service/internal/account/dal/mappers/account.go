package mappers

import (
	"github.com/shortener-service/internal/account/dal/models"
	"github.com/shortener-service/internal/account/domain/entities"
	ValueObjects "github.com/shortener-service/internal/account/domain/value-objects"
)

func MapToAccountModel(e *entities.AccountEntity) *models.AccountModel {
	return &models.AccountModel{
		ID:        e.ID.GetValue(),
		Name:      e.Name,
		Username:  e.Username.GetValue(),
		Password:  e.Password.GetValue(),
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}

func MapToAccountDomain(m *models.AccountModel) *entities.AccountEntity {
	ID, _ := ValueObjects.NewID(m.ID)
	Username, _ := ValueObjects.NewUsername(m.Username)
	Password := ValueObjects.NewHashedPassword(m.Password)

	return &entities.AccountEntity{
		ID:        ID,
		Name:      m.Name,
		Username:  Username,
		Password:  Password,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}
