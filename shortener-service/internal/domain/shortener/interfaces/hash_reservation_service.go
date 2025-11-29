package interfaces

type HashReservationService interface {
	Reserve(length int) string
}
