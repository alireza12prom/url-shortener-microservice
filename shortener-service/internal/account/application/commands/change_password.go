package commands

import (
	"github.com/alireza12prom/url-shortener-microservice/shortener-service/internal/account/application/dto"
	"github.com/alireza12prom/url-shortener-microservice/shortener-service/internal/account/domain/repositories"
)

type ChangePasswordCommand struct {
	AccountRepository *repositories.AccountRepository
}

func NewUpdateAccountCommand(accountRepository *repositories.AccountRepository) *ChangePasswordCommand {
	return &ChangePasswordCommand{
		AccountRepository: accountRepository,
	}
}

func (command *ChangePasswordCommand) Execute(input dto.ChangePasswordInput) (*dto.ChangePasswordOutput, error) {
	return nil, nil
}
