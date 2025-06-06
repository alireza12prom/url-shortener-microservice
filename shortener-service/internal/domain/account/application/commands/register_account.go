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
	ID        string `json:"id"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
