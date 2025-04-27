package ValueObjects

import (
	"golang.org/x/crypto/bcrypt"
)

type Password struct {
	value string
}

func NewPassword(value string) (*Password, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(value), 14)
	if err != nil {
		return nil, err
	}

	return &Password{
		value: string(bytes),
	}, nil
}

func NewHashedPassword(value string) *Password {
	return &Password{
		value: value,
	}
}

func (p *Password) GetValue() string {
	return p.value
}

func (p *Password) Compare(value string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(p.value), []byte(value))
	return err == nil
}
