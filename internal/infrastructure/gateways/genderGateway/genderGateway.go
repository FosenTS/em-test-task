package genderGateway

import (
	"context"
	"em-test-task/internal/application/config"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
)

type GenderGateway interface {
	GetGenderByName(ctx context.Context, name string) (string, error)
}

type genderGateway struct {
	config *config.GenderGatewayConfig

	log *logrus.Entry
}

func NewGenderGateway(config *config.GenderGatewayConfig, log *logrus.Entry) GenderGateway {
	return &genderGateway{config: config, log: log}
}

func (gG *genderGateway) GetGenderByName(ctx context.Context, name string) (string, error) {
	queries := url.Values{}

	queries.Set("name", name)

	url := url.URL{
		Scheme:   "https",
		Host:     gG.config.GenderUrl,
		RawQuery: queries.Encode(),
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url.String(), nil)
	if err != nil {
		gG.log.Errorln(err)
		return "", err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		gG.log.Errorln(err)
		return "", err
	}
	var genderResp *genderResponce
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&genderResp)
	if err != nil {
		gG.log.Errorln(err)
		return "", err
	}

	return genderResp.Gender, nil
}

type genderResponce struct {
	Count       int     `json:"count" binding:"required"`
	Name        string  `json:"name" binding:"required"`
	Gender      string  `json:"gender" binding:"required"`
	Probability float64 `json:"probability" binding:"required"`
}
