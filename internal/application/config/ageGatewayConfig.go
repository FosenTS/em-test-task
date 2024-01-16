package config

import (
	"github.com/spf13/viper"
)

const ageGatewayConfigFilename = "ageGateway.config.yaml"

type AgeGatewayConfig struct {
	AgeUrl string
}

var AgeGatewayInst = &AgeGatewayConfig{}

func AgeGateway() (*AgeGatewayConfig, error) {
	ageGatewayConfigViper := viper.New()
	ageGatewayConfigViper.AddConfigPath(viper.GetString("configPath"))
	ageGatewayConfigViper.SetConfigName(ageGatewayConfigFilename)
	ageGatewayConfigViper.SetConfigType("yaml")
	err := ageGatewayConfigViper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	AgeGatewayInst.AgeUrl = ageGatewayConfigViper.GetString("ageUrl")

	return AgeGatewayInst, nil
}
