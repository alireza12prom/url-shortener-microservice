package commands

type ChangeUsernameCommand struct {
	UserID   string
	Username string
}

func NewChangeUsernameCommand(userId, username string) ChangeUsernameCommand {
	return ChangeUsernameCommand{
		UserID:   userId,
		Username: username,
	}
}

type ChangeUsernameCommandOutput struct{}
