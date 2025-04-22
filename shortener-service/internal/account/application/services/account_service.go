package services

import (
	"errors"

	"github.com/google/uuid"
	"github.com/shortener-service/internal/account/application/commands"
	"github.com/shortener-service/internal/account/application/queries"
	"github.com/shortener-service/internal/account/domain/entities"
	"github.com/shortener-service/internal/account/domain/interfaces"
	"github.com/shortener-service/internal/infrastructure/jwt"
)

type AccountService struct {
	accountRepo interfaces.AccountRepository
}

func NewAccountService(accountRepo interfaces.AccountRepository) *AccountService {
	return &AccountService{
		accountRepo: accountRepo,
	}
}

func (s *AccountService) RegisterAccount(command *commands.RegisterAccountCommand) (
	*commands.RegisterAccountCommandResult,
	error,
) {
	// check "username" is unique
	isUsernameUnique, err := s.accountRepo.IsUsernameUnique(command.Username)
	if err != nil {
		return nil, err
	}

	if !isUsernameUnique {
		return nil, errors.New("username is not unique")
	}

	// check "email" is unique
	isEmailUnique, err := s.accountRepo.IsEmailUnique(command.Email)
	if err != nil {
		return nil, err
	}

	if !isEmailUnique {
		return nil, errors.New("email is not unique")
	}

	// create a new account
	entity, err := entities.NewAccount(
		uuid.New().String(),
		command.Username,
		command.Email,
		command.Password,
	)
	if err != nil {
		return nil, err
	}

	err = s.accountRepo.Save(entity)
	if err != nil {
		return nil, err
	}

	token, err := jwt.GenerateToken(entity.ID.GetValue())
	if err != nil {
		return nil, err
	}

	return &commands.RegisterAccountCommandResult{
		Token: token,
	}, nil
}

func (s *AccountService) ChangePassword(command *commands.ChangePasswordCommand) (
	*commands.ChangePasswordCommandOutput,
	error,
) {
	account, err := s.accountRepo.GetByUserId(command.UserID)
	if err != nil {
		return nil, err
	}

	isPasswordMatch, err := account.Password.Compare(command.OldPassword)
	if err != nil {
		return nil, err
	}

	if !isPasswordMatch {
		return nil, errors.New("'old_password' isn't match")
	}

	account.UpdatePassword(command.NewPassword)

	err = s.accountRepo.Save(account)
	if err != nil {
		return nil, err
	}

	return &commands.ChangePasswordCommandOutput{}, nil
}

func (s *AccountService) ChangeUsername(command *commands.ChangeUsernameCommand) (
	*commands.ChangeUsernameCommandOutput,
	error,
) {
	isUsernameUnique, err := s.accountRepo.IsUsernameUnique(command.Username)
	if err != nil {
		return nil, err
	}

	if !isUsernameUnique {
		return nil, errors.New("username has already taken")
	}

	account, err := s.accountRepo.GetByUserId(command.UserID)
	if err != nil {
		return nil, err
	}

	account.UpdateUsername(command.Username)

	err = s.accountRepo.Save(account)
	if err != nil {
		return nil, err
	}

	return &commands.ChangeUsernameCommandOutput{}, nil
}

func (s *AccountService) GetAccount(query *queries.GetAccountQuery) (
	*queries.GetAccountQueryResult,
	error,
) {
	account, err := s.accountRepo.GetByUserId(query.UserID)
	if err != nil {
		return nil, err
	}

	return &queries.GetAccountQueryResult{
		ID:        account.ID.GetValue(),
		Email:     account.Email.GetValue(),
		Username:  account.Username.GetValue(),
		CreatedAt: account.CreatedAt.String(),
		UpdatedAt: account.UpdatedAt.String(),
	}, nil
}
