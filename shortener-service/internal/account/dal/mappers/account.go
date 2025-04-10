package mappers

import (
	"github.com/alireza12prom/url-shortener-microservice/shortener-service/internal/account/dal/models"
	entity "github.com/alireza12prom/url-shortener-microservice/shortener-service/internal/account/domain/Entity"
	"github.com/alireza12prom/url-shortener-microservice/shortener-service/internal/account/domain/ValueObject"
)

func MapToAccountModel(e *entity.AccountEntity) *models.AccountModel {
	return &models.AccountModel{
		ID:        e.ID.GetValue(),
		Username:  e.Username.GetValue(),
		Email:     e.Email.GetValue(),
		Password:  e.Password.GetValue(),
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}

func MapToAccountDomain(m *models.AccountModel) *entity.AccountEntity {
	ID, _ := ValueObject.NewID(m.ID)
	Username, _ := ValueObject.NewUsername(m.Username)
	Email, _ := ValueObject.NewEmail(m.Email)
	Password, _ := ValueObject.NewPassword(m.Password)

	return &entity.AccountEntity{
		ID:        ID,
		Username:  Username,
		Email:     Email,
		Password:  Password,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}
