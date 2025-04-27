package commands

import "github.com/shortener-service/internal/account/application/dto"

type LoginAccountCommand struct {
	Password string
	Username string
}

func NewLoginAccountCommand(input *dto.LoginAccountInput) LoginAccountCommand {
	return LoginAccountCommand{
		Password: input.Password,
		Username: input.Username,
	}
}

type LoginAccountCommandResult struct {
	Token string `json:"token"`
}
