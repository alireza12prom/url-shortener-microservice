package mappers

import (
	"github.com/shortener-service/internal/account/dal/models"
	"github.com/shortener-service/internal/account/domain/entities"
	ValueObjects "github.com/shortener-service/internal/account/domain/value-objects"
)

func MapToAccountModel(e *entities.AccountEntity) *models.AccountModel {
	return &models.AccountModel{
		ID:        e.ID.GetValue(),
		Username:  e.Username.GetValue(),
		Email:     e.Email.GetValue(),
		Password:  e.Password.GetValue(),
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}

func MapToAccountDomain(m *models.AccountModel) *entities.AccountEntity {
	ID, _ := ValueObjects.NewID(m.ID)
	Username, _ := ValueObjects.NewUsername(m.Username)
	Email, _ := ValueObjects.NewEmail(m.Email)
	Password, _ := ValueObjects.NewPassword(m.Password)

	return &entities.AccountEntity{
		ID:        ID,
		Username:  Username,
		Email:     Email,
		Password:  Password,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}
