package config

import (
	"github.com/spf13/viper"
)

const nationalGatewayConfigFilename = "nationalGateway.config.yaml"

type NationalGatewayConfig struct {
	NationalUrl string
}

var nationalGatewayInst = &NationalGatewayConfig{}

func NationalGateway() (*NationalGatewayConfig, error) {
	nationalGatewayConfigViper := viper.New()
	nationalGatewayConfigViper.AddConfigPath(viper.GetString("configPath"))
	nationalGatewayConfigViper.SetConfigName(nationalGatewayConfigFilename)
	nationalGatewayConfigViper.SetConfigType("yaml")
	err := nationalGatewayConfigViper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	nationalGatewayInst.NationalUrl = nationalGatewayConfigViper.GetString("nationalUrl")

	return nationalGatewayInst, nil
}
