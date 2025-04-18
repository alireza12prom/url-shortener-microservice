package queries

import "github.com/shortener-service/internal/account/application/dto"

type GetAccountQuery struct {
	UserID string
}

func NewGetAccountQuery(input *dto.GetAccountInput) GetAccountQuery {
	return GetAccountQuery{
		UserID: input.UserID,
	}
}

type GetAccountQueryResult struct {
	ID        string
	Username  string
	Email     string
	CreatedAt string
	UpdatedAt string
}
