package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type AppConfig struct {
	LogLevel logrus.Level
}

var appConfigInst = &AppConfig{}

func App() (*AppConfig, error) {
	logLevel := viper.GetString("logLevel")

	if logLevel != "" {
		level, err := logrus.ParseLevel(logLevel)
		if err != nil {
			return nil, err
		}
		appConfigInst.LogLevel = level
	} else {
		appConfigInst.LogLevel = logrus.InfoLevel
	}

	return appConfigInst, nil
}
