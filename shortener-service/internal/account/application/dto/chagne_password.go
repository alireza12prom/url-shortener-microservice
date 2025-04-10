package dto

type ChangePasswordInput struct {
	UserID      string `json:"UserID"`
	NewPassword string `json:"NewPassword"`
	OldPassword string `json:"OldPassword"`
}

type ChangePasswordOutput struct {
}
