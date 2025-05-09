package interfaces

import (
	"github.com/shortener-service/internal/domain/account/application/commands"
	"github.com/shortener-service/internal/domain/account/application/queries"
)

type AccountService interface {
	ChangePassword(command *commands.ChangePasswordCommand) (*commands.ChangePasswordCommandOutput, error)
	RegisterAccount(command *commands.RegisterAccountCommand) (*commands.RegisterAccountCommandResult, error)
	LoginAccount(command *commands.LoginAccountCommand) (*commands.LoginAccountCommandResult, error)
	ChangeUsername(command *commands.ChangeUsernameCommand) (*commands.ChangeUsernameCommandOutput, error)
	GetAccount(query *queries.GetAccountQuery) (*queries.GetAccountQueryResult, error)
}
