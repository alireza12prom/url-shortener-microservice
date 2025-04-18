package commands

import "github.com/shortener-service/internal/account/application/dto"

type RegisterAccountCommand struct {
	Name     string
	Email    string
	Password string
	Username string
}

func NewRegisterAccountCommand(input *dto.RegisterAccountInput) RegisterAccountCommand {
	return RegisterAccountCommand{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
		Username: input.Username,
	}
}

type RegisterAccountCommandResult struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
