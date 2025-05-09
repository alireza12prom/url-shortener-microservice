package ValueObjects

import (
	"errors"
	"net/url"
)

type Endpoint struct {
	value string
}

func NewEndpoint(value string) (*Endpoint, error) {
	u, err := url.ParseRequestURI(value)
	if err != nil {
		return nil, errors.New("invalid endpoint")
	}

	return &Endpoint{value: u.String()}, nil
}

func (s *Endpoint) GetValue() string {
	return s.value
}
