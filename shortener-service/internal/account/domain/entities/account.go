package entities

import (
	"time"

	ValueObjects "github.com/shortener-service/internal/account/domain/value-objects"
)

type AccountEntity struct {
	ID        *ValueObjects.ID
	Name      string
	Username  *ValueObjects.Username
	Password  *ValueObjects.Password
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewAccount(id, name, username, password string) (*AccountEntity, error) {
	ID, err := ValueObjects.NewID(id)
	if err != nil {
		return nil, err
	}

	Username, err := ValueObjects.NewUsername(username)
	if err != nil {
		return nil, err
	}

	Password, err := ValueObjects.NewPassword(password)
	if err != nil {
		return nil, err
	}

	return &AccountEntity{
		ID:        ID,
		Name:      name,
		Username:  Username,
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
