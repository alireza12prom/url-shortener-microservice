package kafka_interface

import (
	"github.com/warehouse-service/internal/common/configs"
	"github.com/warehouse-service/internal/domain/domain-events/application/services"
	"github.com/warehouse-service/internal/domain/domain-events/dal/repositories"
	"github.com/warehouse-service/internal/domain/domain-events/interfaces/kafka/handlers"
	"github.com/warehouse-service/internal/infrastructure/clickhouse"
	"github.com/warehouse-service/internal/infrastructure/kafka"
)

func Setup(ClickHouseDB *clickhouse.ClickHouseDB) {
	DomainEventRepo := repositories.NewDomainEventRepository(ClickHouseDB)

	// Setup Account Event Consumer

	ShortenerServiceAccountEventConsumer := kafka.NewKafkaConsumer(
		&kafka.KafkaConsumerConfig{
			Group: configs.KAFKA_GROUP_ACCOUNT,
			Topic: configs.KAFKA_TOPIC_ACCOUNT,
		},
	)

	kafka.NewKafkaConsumerService(
		ShortenerServiceAccountEventConsumer,
		handlers.NewCaptureEventConsumer(
			&services.CaptureEventService{EventRepo: DomainEventRepo},
		),
	).Start()
}
