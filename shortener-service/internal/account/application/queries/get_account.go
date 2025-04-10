package queries

import (
	"github.com/alireza12prom/url-shortener-microservice/shortener-service/internal/account/application/dto"
	"github.com/alireza12prom/url-shortener-microservice/shortener-service/internal/account/domain/repositories"
)

type GetAccountQuery struct {
	AccountRepository *repositories.AccountRepository
}

func NewGetAccountQuery(accountRepository *repositories.AccountRepository) *GetAccountQuery {
	return &GetAccountQuery{
		AccountRepository: accountRepository,
	}
}

func (query *GetAccountQuery) Execute(input dto.GetAccountInput) (*dto.GetAccountOutput, error) {
	return nil, nil
}
