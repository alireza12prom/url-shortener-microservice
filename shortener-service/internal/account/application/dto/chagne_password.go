package dto

type ChangePasswordInput struct {
	UserID      string `json:"userId" binding:"required,uuid"`
	NewPassword string `json:"newPassword" binding:"required,min=5,max=15,alphanum"`
	OldPassword string `json:"oldPassword" binding:"required,min=5,max=15,alphanum"`
}
