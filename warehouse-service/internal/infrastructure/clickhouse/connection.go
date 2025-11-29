package clickhouse

import (
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/warehouse-service/internal/common/configs"
	"github.com/warehouse-service/internal/common/logger"
)

type ClickHouseDB struct {
	session clickhouse.Conn
	logger  *logger.Logger
}

func NewClickHouseDB() *ClickHouseDB {
	logger := logger.NewLogger("ClickHouse")

	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: configs.CLICK_HOUSE_HOSTS,
		Auth: clickhouse.Auth{
			Database: configs.CLICK_HOUSE_DATABASE,
			Username: configs.CLICK_HOUSE_USER,
			Password: configs.CLICK_HOUSE_PASS,
		},
	})
	if err != nil {
		logger.Error(err.Error())
	}

	logger.Info("Connected")

	return &ClickHouseDB{session: conn, logger: logger}
}

func (db *ClickHouseDB) GetSession() clickhouse.Conn {
	return db.session
}
