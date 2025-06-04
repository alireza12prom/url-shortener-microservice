package scylladb

import (
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/warehouse-service/internal/common/logger"
)

type ClickHouseDB struct {
	session *clickhouse.Conn
	logger  *logger.Logger
}

func NewClickHouseDB() *ClickHouseDB {
	logger := logger.NewLogger("ClickHouse")

	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{"localhost:9000"},
		Auth: clickhouse.Auth{
			Database: "default",
			Username: "default",
			Password: "",
		},
	})
	if err != nil {
		logger.Error(err.Error())
	}

	logger.Info("Connected")

	return &ClickHouseDB{session: &conn, logger: logger}
}
