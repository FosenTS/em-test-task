package product

import (
	httpController "em-test-task/internal/infrastructure/controllers/handler"
	"em-test-task/internal/infrastructure/controllers/handler/bioHandler"
	"github.com/sirupsen/logrus"
)

type Controllers struct {
	httpController httpController.HTTPHandler
}

func NewControllers(gateways *Gateways, log *logrus.Entry) *Controllers {
	return &Controllers{
		httpController: httpController.NewHttpController(
			bioHandler.NewBioHandler(gateways.bioInfoService, gateways.filterService, gateways.AgeGateway, gateways.GenderGateway, gateways.NationalGateway, log.WithField("location", "BioHandler")),
		),
	}
}
