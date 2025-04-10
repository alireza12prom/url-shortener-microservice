package dto

type ChangeUsernameInput struct {
	UserID   string `json:"id"`
	Username string `json:"name"`
}

type ChangeUsernameOutput struct {
	Username string `json:"id"`
}
