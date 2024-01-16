package services

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

type FilterService interface {
	ValidateFilters(raw map[string]string) (map[string]string, map[string]string, error)
}

type filterService struct {
	log *logrus.Entry
}

func NewFilterService(log *logrus.Entry) FilterService {
	return &filterService{log: log}
}

func (fS *filterService) ValidateFilters(raw map[string]string) (map[string]string, map[string]string, error) {
	validFilters := make(map[string]string)
	invalidFiltes := make(map[string]string)
	for name, value := range raw {
		typeName, err := fS.validateType(name)
		if err != nil {
			invalidFiltes[typeName] = value
			fS.log.Debugln("invalid filter type: ", typeName)
			continue
		}
		validFilters[typeName] = value
	}

	return validFilters, invalidFiltes, nil
}

func (fS *filterService) validateType(t string) (string, error) {
	switch t {
	case "name":
		return filterName, nil
	case "surname":
		return filterSurname, nil
	case "age":
		return filterAge, nil
	default:
		return t, undefinedTypeError
	}
}

var undefinedTypeError = fmt.Errorf("undefined filter type")

const (
	filterName    = "name"
	filterSurname = "surname"
	filterAge     = "age"
)
