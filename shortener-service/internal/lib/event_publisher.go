package lib

type EventPublisher interface {
	Publish(event *Event) error
}
