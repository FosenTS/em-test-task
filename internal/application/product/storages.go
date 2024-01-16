package product

import (
	"em-test-task/internal/domain/storages"
	"em-test-task/internal/domain/storages/postgres"
	"em-test-task/pkg/db/postgresql"
	"github.com/sirupsen/logrus"
)

type Storages struct {
	BioInfo storages.BioInfoStorage
}

func NewStorages(pool postgresql.PGXPool, log *logrus.Entry) *Storages {
	return &Storages{
		BioInfo: postgres.NewBioInfoRepository(pool, log.WithField("location", "bioInfoRepository")),
	}
}
