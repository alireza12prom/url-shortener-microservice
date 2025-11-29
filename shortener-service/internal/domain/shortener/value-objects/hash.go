package valueobjects

import (
	"errors"
	"regexp"
)

var HashPattern = regexp.MustCompile(`^[a-zA-Z0-9]*$`)

type Hash struct {
	value string
}

func NewHash(value string) (*Hash, error) {
	if len(value) < 5 {
		return nil, errors.New("invalid hash length")
	}

	if !HashPattern.MatchString(value) {
		return nil, errors.New("invalid hash format")
	}

	return &Hash{value: value}, nil
}

func (s *Hash) GetValue() string {
	return s.value
}
