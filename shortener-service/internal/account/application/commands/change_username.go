package commands

import (
	"github.com/shortener-service/internal/account/application/dto"
	"github.com/shortener-service/internal/account/domain/repositories"
)

type ChangeUsernameCommand struct {
	AccountRepository *repositories.AccountRepository
}

func NewChangeUsernameCommand(accountRepository *repositories.AccountRepository) *ChangeUsernameCommand {
	return &ChangeUsernameCommand{
		AccountRepository: accountRepository,
	}
}

func (command *ChangeUsernameCommand) Execute(input dto.ChangeUsernameInput) (*dto.ChangeUsernameOutput, error) {
	return nil, nil
}
