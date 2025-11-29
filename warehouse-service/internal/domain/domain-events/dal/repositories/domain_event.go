package repositories

import (
	"context"

	ch "github.com/ClickHouse/clickhouse-go/v2"
	"github.com/warehouse-service/internal/common/exceptions"
	"github.com/warehouse-service/internal/domain/domain-events/dal/mappers"
	"github.com/warehouse-service/internal/domain/domain-events/entities"
	"github.com/warehouse-service/internal/infrastructure/clickhouse"
)

type DomainEventRepository struct {
	session ch.Conn
}

func NewDomainEventRepository(connection *clickhouse.ClickHouseDB) *DomainEventRepository {
	return &DomainEventRepository{
		session: connection.GetSession(),
	}
}

func (r *DomainEventRepository) Save(entity *entities.DomainEventEntity) error {
	query := "INSERT INTO domain_events (id, name, context, payload, datetime, correlation_id) VALUES (?, ?, ?, ?, ?, ?)"

	model := mappers.MapToAccountModel(entity)

	err := r.session.Exec(
		context.Background(),
		query,
		model.ID,
		model.Name,
		model.Context,
		model.Payload,
		model.DateTime,
		model.CorrelationID,
	)
	if err != nil {
		return exceptions.NewDatabaseException(err)
	}

	return nil
}
