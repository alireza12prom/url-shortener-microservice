package kafka

import (
	"context"
	"encoding/json"

	"github.com/segmentio/kafka-go"
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
}

func NewKafkaPublisher(config *KafkaPublisherConfig) *KafkaPublisher {
	conn, err := kafka.Dial("tcp", config.Addrs[0])
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	err = conn.CreateTopics(kafka.TopicConfig{
		Topic:             config.Topic.Name,
		NumPartitions:     config.Topic.Partitions,
		ReplicationFactor: 1,
	})
	if err != nil {
		if err.Error() != "Topic already exists" {
			panic(err.Error())
		}
	}

	return &KafkaPublisher{
		writer: &kafka.Writer{
			Addr:     kafka.TCP(config.Addrs...),
			Balancer: &kafka.LeastBytes{},
		},
		config: config,
	}
}

func (Self *KafkaPublisher) Publish(event *lib.Event) error {
	payload, err := json.Marshal(event)
	if err != nil {
		return err
	}

	msg := kafka.Message{
		Topic: Self.config.Topic.Name,
		Key:   []byte(event.AggregateID()),
		Value: payload,
	}

	return Self.writer.WriteMessages(context.Background(), msg)
}
