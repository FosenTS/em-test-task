package product

import (
	"em-test-task/internal/domain/services"
	"github.com/sirupsen/logrus"
)

type Services struct {
	bioInfoService services.BioInfoService
	filterService  services.FilterService
}

func NewServices(storages *Storages, log *logrus.Entry) *Services {
	return &Services{
		bioInfoService: services.NewBioInfoService(storages.BioInfo, log.WithField("location", "bioInfoService")),
		filterService:  services.NewFilterService(log.WithField("location", "filterService")),
	}
}
