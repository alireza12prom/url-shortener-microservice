package kafka

import (
	"github.com/warehouse-service/internal/lib"
)

type KafkaConsumerService struct {
	consumer lib.KafkaConsumer
	handler  lib.KafkaHandler
}

func NewKafkaConsumerService(consumer lib.KafkaConsumer, handler lib.KafkaHandler) *KafkaConsumerService {
	return &KafkaConsumerService{consumer: consumer, handler: handler}
}

func (s *KafkaConsumerService) Start() error {
	return s.consumer.Consume(s.handler.Handle)
}
