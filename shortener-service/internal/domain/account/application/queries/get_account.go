package queries

import "github.com/shortener-service/internal/domain/account/application/dto"

type GetAccountQuery struct {
	UserID string
}

func NewGetAccountQuery(input *dto.GetAccountInput) GetAccountQuery {
	return GetAccountQuery{
		UserID: input.UserID,
	}
}

type GetAccountQueryResult struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
