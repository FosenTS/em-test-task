package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type HTTPConfig struct {
	httpHost   string
	httpPort   int
	HttpServer string
}

var httpConfig = &HTTPConfig{}

func Http() (*HTTPConfig, error) {
	httpConfig.httpHost = viper.GetString("httpHost")
	httpConfig.httpPort = viper.GetInt("httpPort")

	httpConfig.HttpServer = fmt.Sprintf("%s:%d", httpConfig.httpHost, httpConfig.httpPort)

	return httpConfig, nil
}
