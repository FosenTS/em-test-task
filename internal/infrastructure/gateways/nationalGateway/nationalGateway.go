package nationalGateway

import (
	"context"
	"em-test-task/internal/application/config"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
)

type NationalGateway interface {
	GetNationalByName(ctx context.Context, name string) (string, error)
}

type nationalGateway struct {
	config *config.NationalGatewayConfig

	log *logrus.Entry
}

func NewNationalGateway(config *config.NationalGatewayConfig, log *logrus.Entry) NationalGateway {
	return &nationalGateway{config: config, log: log}
}

func (nG *nationalGateway) GetNationalByName(ctx context.Context, name string) (string, error) {
	queries := url.Values{}

	queries.Set("name", name)

	url := url.URL{
		Scheme:   "https",
		Host:     nG.config.NationalUrl,
		RawQuery: queries.Encode(),
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url.String(), nil)
	if err != nil {
		nG.log.Errorln(err)
		return "", err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		nG.log.Errorln(err)
		return "", err
	}
	var nationalResp *nationalResponce
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&nationalResp)
	if err != nil {
		nG.log.Errorln(err)
		return "", err
	}
	if nationalResp.Country == nil {
		err = fmt.Errorf("empty country array")
		nG.log.Errorln(err)
		return "", err
	}
	type upCountry struct {
		id          string
		probability float64
	}
	var up = upCountry{
		id:          nationalResp.Country[0].CountryId,
		probability: nationalResp.Country[0].Probability,
	}
	for _, country := range nationalResp.Country {
		if up.probability < country.Probability {
			up.id = country.CountryId
			up.probability = country.Probability
		}
	}

	return up.id, nil
}

type nationalResponce struct {
	Count   int    `json:"count" binding:"required"`
	Name    string `json:"name" binding:"required"`
	Country []struct {
		CountryId   string  `json:"country_id" binding:"required"`
		Probability float64 `json:"probability" binding:"required"`
	}
}
