package commands

import "github.com/shortener-service/internal/account/application/dto"

type ChangeUsernameCommand struct {
	UserID   string
	Username string
}

func NewChangeUsernameCommand(input *dto.ChangeUsernameInput) ChangeUsernameCommand {
	return ChangeUsernameCommand{
		UserID:   input.UserID,
		Username: input.Username,
	}
}

type ChangeUsernameCommandOutput struct{}
