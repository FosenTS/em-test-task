package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

const envFileName = ".env"

func Env() error {
	err := godotenv.Load(envFileName)
	if err != nil {
		return err
	}

	err = viper.BindEnv("configPath", "CONFIG_PATH")
	if err != nil {
		return err
	}

	err = viper.BindEnv("logLevel", "LOG_LEVEL")

	err = viper.BindEnv("httpHost", "HTTP_HOST")
	if err != nil {
		return err
	}

	err = viper.BindEnv("httpPort", "HTTP_PORT")
	if err != nil {
		return err
	}

	err = viper.BindEnv("hostDB", "HOST_DB")
	if err != nil {
		return err
	}
	err = viper.BindEnv("portDB", "PORT_DB")
	if err != nil {
		return err
	}
	err = viper.BindEnv("usernameDB", "USERNAME_DB")
	if err != nil {
		return err
	}
	err = viper.BindEnv("passwordDB", "PASSWORD_DB")
	if err != nil {
		return err
	}
	err = viper.BindEnv("dbName", "DB_NAME")
	if err != nil {
		return err
	}
	err = viper.BindEnv("sslMode", "SSL_MODE")
	if err != nil {
		return err
	}

	return nil
}
