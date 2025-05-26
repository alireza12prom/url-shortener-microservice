package kafka

import (
	"context"
	"encoding/json"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/shortener-service/internal/common/logger"
	"github.com/shortener-service/internal/lib"
)

type (
	Topic struct {
		Name       string
		Partitions int
	}

	KafkaPublisherConfig struct {
		Addrs []string
		Topic Topic
	}
)

type KafkaPublisher struct {
	writer *kafka.Writer
	config *KafkaPublisherConfig
	logger *logger.Logger
}

func NewKafkaPublisher(config *KafkaPublisherConfig) *KafkaPublisher {
	return &KafkaPublisher{
		writer: &kafka.Writer{
			Addr:                   kafka.TCP(config.Addrs...),
			Balancer:               &kafka.LeastBytes{},
			AllowAutoTopicCreation: true,
			Topic:                  config.Topic.Name,
		},
		config: config,
		logger: logger.NewLogger("KafkaPublisher"),
	}
}

func (Self *KafkaPublisher) Publish(event *lib.Event) error {
	payload, err := json.Marshal(event)
	if err != nil {
		return err
	}

	msg := kafka.Message{
		Value: payload,
		Time:  time.Now(),
	}

	Self.logger.Debug(
		"Publish new event",
		logger.Fields{"topic": Self.config.Topic.Name, "event": event},
	)

	return Self.writer.WriteMessages(context.Background(), msg)
}
