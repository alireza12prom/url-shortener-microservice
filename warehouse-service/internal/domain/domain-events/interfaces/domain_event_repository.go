package interfaces

import "github.com/warehouse-service/internal/domain/domain-events/entities"

type DomainEventRepository interface {
	Save(entity *entities.DomainEventEntity) error
}
