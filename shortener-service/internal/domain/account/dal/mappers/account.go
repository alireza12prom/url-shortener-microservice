package mappers

import (
	"github.com/shortener-service/internal/domain/account/dal/models"
	"github.com/shortener-service/internal/domain/account/entities"
	"github.com/shortener-service/internal/domain/account/value-objects"
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
	ID, _ := valueobjects.NewID(m.ID)
	Username, _ := valueobjects.NewUsername(m.Username)
	Password := valueobjects.NewHashedPassword(m.Password)

	return &entities.AccountEntity{
		ID:        ID,
		Name:      m.Name,
		Username:  Username,
		Password:  Password,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}
