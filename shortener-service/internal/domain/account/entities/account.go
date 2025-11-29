package entities

import (
	"time"

	"github.com/shortener-service/internal/domain/account/value-objects"
)

type AccountEntity struct {
	ID        *valueobjects.ID
	Name      string
	Username  *valueobjects.Username
	Password  *valueobjects.Password
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewAccount(id, name, username, password string) (*AccountEntity, error) {
	ID, err := valueobjects.NewID(id)
	if err != nil {
		return nil, err
	}

	Username, err := valueobjects.NewUsername(username)
	if err != nil {
		return nil, err
	}

	Password, err := valueobjects.NewPassword(password)
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

func (a *AccountEntity) UpdateUsername(value string) error {
	username, err := valueobjects.NewUsername(value)
	if err != nil {
		return err
	}

	a.Username = username
	a.UpdatedAt = time.Now()

	return nil
}

func (a *AccountEntity) UpdatePassword(value string) error {
	password, err := valueobjects.NewPassword(value)
	if err != nil {
		return err
	}

	a.Password = password
	a.UpdatedAt = time.Now()

	return nil
}
