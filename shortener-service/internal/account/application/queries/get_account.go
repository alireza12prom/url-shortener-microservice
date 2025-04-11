package queries

import (
	"github.com/shortener-service/internal/account/application/dto"
	"github.com/shortener-service/internal/account/domain/repositories"
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
