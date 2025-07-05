package lib

type KafkaConsumer interface {
	Consume(handler func([]byte) error) error
}
