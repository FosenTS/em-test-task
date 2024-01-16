package services

import (
	"em-test-task/internal/domain/entity"
	"em-test-task/internal/domain/storages"
	"em-test-task/internal/domain/storages/dto"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

type BioInfoService interface {
	Add(ctx context.Context, bIC *dto.BioInfoCreate) (*entity.BioInfo, error)
	DeleteById(ctx context.Context, id int) (*entity.BioInfo, error)
	UpdateById(ctx context.Context, info *entity.BioInfo) error
	GetAll(ctx context.Context) ([]*entity.BioInfo, error)
	GetByFilters(ctx context.Context, filters map[string]string) ([]*entity.BioInfo, error)
}

type bioService struct {
	bioStorage storages.BioInfoStorage

	log *logrus.Entry
}

func (b bioService) Add(ctx context.Context, bIC *dto.BioInfoCreate) (*entity.BioInfo, error) {
	bI, err := b.bioStorage.Store(ctx, bIC)
	if err != nil {
		return nil, err
	}

	return bI, nil
}

func (b bioService) DeleteById(ctx context.Context, id int) (*entity.BioInfo, error) {
	bI, err := b.bioStorage.DeleteByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return bI, nil
}

func (b bioService) UpdateById(ctx context.Context, info *entity.BioInfo) error {
	return b.bioStorage.UpdateByID(ctx, info)
}

func (b bioService) GetAll(ctx context.Context) ([]*entity.BioInfo, error) {
	bIs, err := b.bioStorage.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return bIs, nil
}

func (b bioService) GetByFilters(ctx context.Context, filters map[string]string) ([]*entity.BioInfo, error) {
	bIs, err := b.bioStorage.GetByFilters(ctx, filters)
	if err != nil {
		return nil, err
	}

	return bIs, nil
}

func NewBioInfoService(bioStorage storages.BioInfoStorage, log *logrus.Entry) BioInfoService {
	return &bioService{bioStorage: bioStorage, log: log}
}
