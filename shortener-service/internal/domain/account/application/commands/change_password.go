package commands

type ChangePasswordCommand struct {
	UserID      string
	NewPassword string
	OldPassword string
}

func NewChangePasswordCommand(userId, newPassword, oldPassword string) ChangePasswordCommand {
	return ChangePasswordCommand{
		UserID:      userId,
		NewPassword: newPassword,
		OldPassword: oldPassword,
	}
}

type ChangePasswordCommandOutput struct{}
