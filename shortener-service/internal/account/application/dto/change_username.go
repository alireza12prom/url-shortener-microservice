package dto

type ChangeUsernameInput struct {
	Username string `json:"username" binding:"required,min=3,max=15,alphanum"`
}
