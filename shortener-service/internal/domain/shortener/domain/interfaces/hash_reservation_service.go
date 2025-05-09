package interfaces

import ValueObjects "github.com/shortener-service/internal/domain/shortener/domain/value-objects"

type HashReservationService interface {
	Reserve() (*ValueObjects.Hash, error)
}
