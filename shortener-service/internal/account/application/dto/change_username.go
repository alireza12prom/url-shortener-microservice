package dto

type ChangeUsernameInput struct {
	UserID   string `json:"userId" binding:"required,uuid"`
	Username string `json:"username" binding:"required,min=3,max=15,alphanum"`
}
