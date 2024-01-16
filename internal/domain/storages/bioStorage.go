package storages

import (
	"context"
	"em-test-task/internal/domain/entity"
	"em-test-task/internal/domain/storages/dto"
)

type BioInfoStorage interface {
	Storage[entity.BioInfo, dto.BioInfoCreate]
	GetByFilters(ctx context.Context, filters map[string]string) ([]*entity.BioInfo, error)
}
