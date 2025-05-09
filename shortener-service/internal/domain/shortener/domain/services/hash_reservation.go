package domain_services

import (
	"errors"

	ValueObjects "github.com/shortener-service/internal/domain/shortener/domain/value-objects"
)

type HashReservationService struct {
	charset string
}

func NewHashReservationService() *HashReservationService {
	return &HashReservationService{
		charset: "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz",
	}
}

func (s *HashReservationService) Reserve(length int) (*ValueObjects.Hash, error) {
	if length <= 0 {
		return nil, errors.New("invalid hash length")
	}

	result := ""
	for length > 0 {
		result = string(s.charset[length%len(s.charset)]) + result
		length /= len(s.charset)
	}

	hash, err := ValueObjects.NewHash(result)
	if err != nil {
		return nil, err
	}

	return hash, nil
}
