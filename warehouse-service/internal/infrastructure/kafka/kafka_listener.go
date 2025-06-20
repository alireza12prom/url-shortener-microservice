package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type (
	KafkaListenerConfig struct {
		Addrs []string
		Group string
		Topic string
	}
)

type KafkaListener struct {
	reader *kafka.Reader
	config *KafkaListenerConfig
}

func NewKafkaListener(config *KafkaListenerConfig) *KafkaListener {
	return &KafkaListener{
		reader: kafka.NewReader(
			kafka.ReaderConfig{
				Brokers:        config.Addrs,
				GroupID:        config.Group,
				Topic:          config.Topic,
				IsolationLevel: kafka.ReadCommitted,
			},
		),
		config: config,
	}
}

func (Self *KafkaListener) Listen(handler func([]byte) error) error {
	println("Starting Kafka listener on topic:", Self.config.Topic)
	ctx := context.Background()

	for {
		m, err := Self.reader.FetchMessage(ctx)
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
