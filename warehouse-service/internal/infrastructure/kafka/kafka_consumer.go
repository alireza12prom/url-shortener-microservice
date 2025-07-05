package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
	"github.com/warehouse-service/internal/common/configs"
)

type (
	KafkaConsumerConfig struct {
		Group string
		Topic string
	}
	KafkaListener struct {
		reader *kafka.Reader
		config *KafkaConsumerConfig
	}
)

func NewKafkaConsumer(config *KafkaConsumerConfig) *KafkaListener {
	return &KafkaListener{
		reader: kafka.NewReader(
			kafka.ReaderConfig{
				Brokers:        configs.KAFKA_BROKERS,
				GroupID:        config.Group,
				Topic:          config.Topic,
				IsolationLevel: kafka.ReadCommitted,
			},
		),
		config: config,
	}
}

func (Self *KafkaListener) Consume(handler func([]byte) error) error {
	println("Starting Kafka listener on topic:", Self.config.Topic)
	ctx := context.Background()

	for {
		m, err := Self.reader.ReadMessage(ctx)
		if err != nil {
			println("Error fetching message:", err.Error())
			return err
		}

		if err := handler(m.Value); err != nil {
			println("Handler error:", err.Error())
		}

		if err := Self.reader.CommitMessages(ctx, m); err != nil {
			println("Failed to commit message:", err)
		}
	}
}
