package dto

type ChangePasswordInput struct {
	NewPassword string `json:"newPassword" binding:"required,min=5,max=15,alphanum"`
	OldPassword string `json:"oldPassword" binding:"required,min=5,max=15,alphanum"`
}
