package ageGateway

import (
	"context"
	"em-test-task/internal/application/config"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
)

type AgeGateway interface {
	GetAgeByName(ctx context.Context, name string) (int, error)
}

type ageGateway struct {
	config *config.AgeGatewayConfig

	log *logrus.Entry
}

func NewAgeGateway(config *config.AgeGatewayConfig, log *logrus.Entry) AgeGateway {
	return &ageGateway{config: config, log: log}
}

func (aG *ageGateway) GetAgeByName(ctx context.Context, name string) (int, error) {
	queries := url.Values{}

	queries.Set("name", name)

	url := url.URL{
		Scheme:   "https",
		Host:     aG.config.AgeUrl,
		RawQuery: queries.Encode(),
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url.String(), nil)
	if err != nil {
		aG.log.Errorln(err)
		return 0, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		aG.log.Errorln(err)
		return 0, err
	}
	var ageResp *ageResponce
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&ageResp)
	if err != nil {
		aG.log.Errorln(err)
		return 0, err
	}

	return ageResp.Age, nil
}

type ageResponce struct {
	Count int    `json:"count" binding:"required"`
	Name  string `json:"name" binding:"required"`
	Age   int    `json:"age" binding:"required"`
}
