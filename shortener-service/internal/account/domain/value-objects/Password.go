package ValueObjects

import "strings"

type Password struct {
	value string
}

func NewPassword(value string) (*Password, error) {
	// TODO: Check if password has the requirement

	return &Password{
		value: strings.ToLower(value),
	}, nil
}

func (id *Password) GetValue() string {
	return id.value
}

func (password *Password) CreateHash() (string, error) {
	// TODO: Implement
	return "xxxx", nil
}

func (password *Password) Compare(value string) (bool, error) {
	// TODO: Implement
	return true, nil
}
