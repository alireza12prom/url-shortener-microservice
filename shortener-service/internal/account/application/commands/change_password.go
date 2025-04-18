package commands

import "github.com/shortener-service/internal/account/application/dto"

type ChangePasswordCommand struct {
	UserID      string
	NewPassword string
	OldPassword string
}

func NewChangePasswordCommand(input *dto.ChangePasswordInput) ChangePasswordCommand {
	return ChangePasswordCommand{
		UserID:      input.UserID,
		NewPassword: input.NewPassword,
		OldPassword: input.OldPassword,
	}
}

type ChangePasswordCommandOutput struct{}
