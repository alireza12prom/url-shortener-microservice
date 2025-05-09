package domain_services

import "math/rand"

type HashReservationService struct {
	charset string
}

func NewHashReservationService() *HashReservationService {
	return &HashReservationService{
		charset: "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz",
	}
}

func (s *HashReservationService) Reserve(length int) string {
	result := make([]byte, length)
	for i := range result {
		result[i] = s.charset[rand.Intn(len(s.charset))]
	}
	return string(result)
}
