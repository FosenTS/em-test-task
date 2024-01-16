package product

import (
	"em-test-task/internal/application/config"
	"em-test-task/internal/infrastructure/gateways/ageGateway"
	"em-test-task/internal/infrastructure/gateways/genderGateway"
	"em-test-task/internal/infrastructure/gateways/nationalGateway"
	"github.com/sirupsen/logrus"
)

type Gateways struct {
	*Services
	ageGateway.AgeGateway
	genderGateway.GenderGateway
	nationalGateway.NationalGateway
}

func NewGateways(services *Services, log *logrus.Entry) (*Gateways, error) {
	ageGatewayCfg, err := config.AgeGateway()
	if err != nil {
		log.Errorln(err)
		return nil, err
	}
	genderGatewayCfg, err := config.GenderGateway()
	if err != nil {
		log.Errorln(err)
		return nil, err
	}
	nationalGatewayCfg, err := config.NationalGateway()
	if err != nil {
		log.Errorln(err)
		return nil, err
	}
	return &Gateways{
		Services:        services,
		AgeGateway:      ageGateway.NewAgeGateway(ageGatewayCfg, log.WithField("location", "ageGateway")),
		GenderGateway:   genderGateway.NewGenderGateway(genderGatewayCfg, log.WithField("location", "genderGateway")),
		NationalGateway: nationalGateway.NewNationalGateway(nationalGatewayCfg, log.WithField("location", "nationalGateway")),
	}, nil
}
