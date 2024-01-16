package storages

import (
	"context"
	"em-test-task/internal/domain/entity"
	"em-test-task/internal/domain/storages/dto"
)

type Storage[Entity entity.Entity, Create dto.Create] interface {
	GetByID(ctx context.Context, id int) (*Entity, error)
	GetAll(ctx context.Context) ([]*Entity, error)
	Store(ctx context.Context, c *Create) (*Entity, error)
	UpdateByID(ctx context.Context, f *Entity) error
	DeleteByID(ctx context.Context, id int) (*Entity, error)
}
