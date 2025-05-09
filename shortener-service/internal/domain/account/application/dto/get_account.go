package dto

type GetAccountInput struct {
	UserID string `json:"userId" binding:"required,uuid"`
}
