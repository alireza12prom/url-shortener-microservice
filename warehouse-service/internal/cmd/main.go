package main

import (
	"github.com/warehouse-service/internal/common/configs"
	kafka_interface "github.com/warehouse-service/internal/domain/domain-events/interfaces/kafka"
	"github.com/warehouse-service/internal/infrastructure/clickhouse"
)

func main() {
	configs.Load()

	// Initialize ClickHouse connection
	clickhouseDB := clickhouse.NewClickHouseDB()

	kafka_interface.Setup(clickhouseDB)
}
