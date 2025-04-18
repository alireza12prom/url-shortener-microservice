package services

import (
	"github.com/shortener-service/internal/account/application/commands"
	"github.com/shortener-service/internal/account/application/queries"
	"github.com/shortener-service/internal/account/domain/interfaces"
)

type AccountService struct {
	repository interfaces.AccountRepository
}

func NewAccountService(repository interfaces.AccountRepository) *AccountService {
	return &AccountService{
		repository: repository,
	}
}

func (s *AccountService) RegisterAccount(command *commands.RegisterAccountCommand) (*commands.RegisterAccountCommandResult, error) {
	return nil, nil
}

func (s *AccountService) ChangePassword(command *commands.ChangePasswordCommand) (*commands.ChangePasswordCommandOutput, error) {
	return nil, nil
}

func (s *AccountService) ChangeUsername(command *commands.ChangeUsernameCommand) (*commands.ChangeUsernameCommandOutput, error) {
	return nil, nil
}

func (s *AccountService) GetAccount(query *queries.GetAccountQuery) (*queries.GetAccountQueryResult, error) {
	return nil, nil
}
