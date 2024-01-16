package config

import (
	"github.com/spf13/viper"
)

const genderGatewayConfigFilename = "genderGateway.config.yaml"

type GenderGatewayConfig struct {
	GenderUrl string
}

var genderGatewayInst = &GenderGatewayConfig{}

func GenderGateway() (*GenderGatewayConfig, error) {
	genderGatewayConfigViper := viper.New()
	genderGatewayConfigViper.AddConfigPath(viper.GetString("configPath"))
	genderGatewayConfigViper.SetConfigName(genderGatewayConfigFilename)
	genderGatewayConfigViper.SetConfigType("yaml")
	err := genderGatewayConfigViper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	genderGatewayInst.GenderUrl = genderGatewayConfigViper.GetString("genderUrl")

	return genderGatewayInst, nil
}
