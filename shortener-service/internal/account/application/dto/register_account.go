package dto

type RegisterAccountInput struct {
	Name     string `json:"name" binding:"required,min=3"`
	Username string `json:"username" binding:"required,min=3,max=15,alphanum"`
	Password string `json:"password" binding:"required,min=5,max=15,alphanum"`
}
