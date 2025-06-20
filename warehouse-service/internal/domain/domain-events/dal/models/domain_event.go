package models

type DomainEventModel struct {
	ID            string
	Name          string
	Context       string
	Payload       string
	DateTime      string
	CorrelationID string
}
