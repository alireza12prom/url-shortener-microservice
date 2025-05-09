package commands

import "github.com/shortener-service/internal/domain/account/application/dto"

type RegisterAccountCommand struct {
	Name     string
	Password string
	Username string
}

func NewRegisterAccountCommand(input *dto.RegisterAccountInput) RegisterAccountCommand {
	return RegisterAccountCommand{
		Name:     input.Name,
		Password: input.Password,
		Username: input.Username,
	}
}

type RegisterAccountCommandResult struct {
	ID        string
	Name      string
	Username  string
	CreatedAt string
	UpdatedAt string
}
