package commands

import (
	"github.com/alireza12prom/url-shortener-microservice/shortener-service/internal/account/application/dto"
	"github.com/alireza12prom/url-shortener-microservice/shortener-service/internal/account/domain/repositories"
)

type RegisterAccountCommand struct {
	AccountRepository *repositories.AccountRepository
}

func NewRegisterAccountCommand(accountRepository *repositories.AccountRepository) *RegisterAccountCommand {
	return &RegisterAccountCommand{
		AccountRepository: accountRepository,
	}
}

func (command *RegisterAccountCommand) Execute(input dto.RegisterAccountInput) (*dto.RegisterAccountOutput, error) {
	return nil, nil
}
