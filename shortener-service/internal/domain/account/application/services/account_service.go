package services

import (
	"log"
	"time"

	"github.com/google/uuid"
	exceptions "github.com/shortener-service/internal/common/exceptions"
	"github.com/shortener-service/internal/domain/account/application/commands"
	"github.com/shortener-service/internal/domain/account/application/queries"
	"github.com/shortener-service/internal/domain/account/entities"
	"github.com/shortener-service/internal/domain/account/events"
	"github.com/shortener-service/internal/domain/account/interfaces"
	"github.com/shortener-service/internal/infrastructure/jwt"
	"github.com/shortener-service/internal/lib"
)

type AccountService struct {
	accountRepo    interfaces.AccountRepository
	eventPublisher lib.EventPublisher
}

func NewAccountService(
	accountRepo interfaces.AccountRepository,
	eventPublisher lib.EventPublisher,
) *AccountService {
	return &AccountService{
		accountRepo:    accountRepo,
		eventPublisher: eventPublisher,
	}
}

func (s *AccountService) RegisterAccount(command *commands.RegisterAccountCommand) (
	*commands.RegisterAccountCommandResult,
	error,
) {
	isUsernameUnique, err := s.accountRepo.IsUsernameUnique(command.Username)
	if err != nil {
		return nil, err
	}

	if !isUsernameUnique {
		return nil, exceptions.NewBusinessException(exceptions.UsernameAlreadyTaken, nil)
	}

	entity, err := entities.NewAccount(
		uuid.New().String(),
		command.Name,
		command.Username,
		command.Password,
	)
	if err != nil {
		return nil, err
	}

	err = s.accountRepo.Save(entity)
	if err != nil {
		return nil, err
	}

	err = s.eventPublisher.Publish(events.NewAccountCreatedEvent(entity))
	if err != nil {
		log.Println("Error publishing event:", err)
	}

	return &commands.RegisterAccountCommandResult{
		ID:        entity.ID.GetValue(),
		Name:      entity.Name,
		Username:  entity.Username.GetValue(),
		CreatedAt: entity.CreatedAt.UTC().Format(time.RFC3339),
		UpdatedAt: entity.UpdatedAt.UTC().Format(time.RFC3339),
	}, nil
}

func (s *AccountService) LoginAccount(command *commands.LoginAccountCommand) (
	*commands.LoginAccountCommandResult,
	error,
) {
	account, err := s.accountRepo.GetByUsername(command.Username)
	if err != nil {
		return nil, err
	}

	if account == nil {
		return nil, exceptions.NewBusinessException(exceptions.WrongAuthenticationPack, nil)
	}

	isPasswordMatch := account.Password.Compare(command.Password)
	if !isPasswordMatch {
		return nil, exceptions.NewBusinessException(exceptions.WrongAuthenticationPack, nil)
	}

	token, err := jwt.GenerateToken(account.ID.GetValue())
	if err != nil {
		return nil, exceptions.NewBusinessException(exceptions.LoginFailed, nil)
	}

	err = s.eventPublisher.Publish(events.NewLoginHappenedEvent(account))
	if err != nil {
		log.Println("Error publishing event:", err)
	}

	return &commands.LoginAccountCommandResult{
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

	isPasswordMatch := account.Password.Compare(command.OldPassword)
	if !isPasswordMatch {
		return nil, exceptions.NewBusinessException(exceptions.WrongPasswordProvided, nil)
	}

	account.UpdatePassword(command.NewPassword)

	err = s.accountRepo.Save(account)
	if err != nil {
		return nil, err
	}

	err = s.eventPublisher.Publish(events.NewPasswordChangedEvent(account))
	if err != nil {
		log.Println("Error publishing event:", err)
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
		return nil, exceptions.NewBusinessException(exceptions.UsernameAlreadyTaken, nil)
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

	err = s.eventPublisher.Publish(events.NewUsernameChangedEvent(account))
	if err != nil {
		log.Println("Error publishing event:", err)
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
		Username:  account.Username.GetValue(),
		CreatedAt: account.CreatedAt.String(),
		UpdatedAt: account.UpdatedAt.String(),
	}, nil
}
