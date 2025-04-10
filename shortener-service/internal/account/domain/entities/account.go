package entities

import (
	"time"

	ValueObjects "github.com/alireza12prom/url-shortener-microservice/shortener-service/internal/account/domain/value-objects"
)

type AccountEntity struct {
	ID        *ValueObjects.ID
	Username  *ValueObjects.Username
	Email     *ValueObjects.Email
	Password  *ValueObjects.Password
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewAccount(id, username, email, password string) (*AccountEntity, error) {
	ID, err := ValueObjects.NewID(id)
	if err != nil {
		return nil, err
	}

	Username, err := ValueObjects.NewUsername(username)
	if err != nil {
		return nil, err
	}

	Email, err := ValueObjects.NewEmail(email)
	if err != nil {
		return nil, err
	}

	Password, err := ValueObjects.NewPassword(password)
	if err != nil {
		return nil, err
	}

	return &AccountEntity{
		ID:        ID,
		Username:  Username,
		Email:     Email,
		Password:  Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (u *AccountEntity) UpdateUsername(value string) error {
	username, err := ValueObjects.NewUsername(value)
	if err != nil {
		return err
	}

	u.Username = username
	u.UpdatedAt = time.Now()

	return nil
}

func (u *AccountEntity) UpdatePassword(value string) error {
	password, err := ValueObjects.NewPassword(value)
	if err != nil {
		return err
	}

	u.Password = password
	u.UpdatedAt = time.Now()
	return nil
}
