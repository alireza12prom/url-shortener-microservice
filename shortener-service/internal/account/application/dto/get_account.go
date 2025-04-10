package dto

type GetAccountInput struct {
	UserID string `json:"userId"`
}

type GetAccountOutput struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
