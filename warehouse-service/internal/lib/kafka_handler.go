package lib

type KafkaHandler interface {
	Handle(input []byte) error
}
