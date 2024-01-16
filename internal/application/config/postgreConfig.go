package config

import (
	"github.com/spf13/viper"
)

const postgreConfigFileName = "postgre.config.yaml"

type PostgeConfig struct {
	Host         string
	Port         int
	User         string
	Password     string
	DatabaseName string
	SSLMode      bool
}

var postgreConfigInst = &PostgeConfig{}

func Postgre() (*PostgeConfig, error) {
	postgreConfigInst.Host = viper.GetString("hostDB")
	postgreConfigInst.Port = viper.GetInt("portDB")
	postgreConfigInst.User = viper.GetString("usernameDB")
	postgreConfigInst.Password = viper.GetString("passwordDB")
	postgreConfigInst.DatabaseName = viper.GetString("dbName")
	postgreConfigInst.SSLMode = viper.GetBool("sslMode")

	postgreConfigViper := viper.New()
	postgreConfigViper.AddConfigPath(viper.GetString("configPath"))
	postgreConfigViper.SetConfigName(postgreConfigFileName)
	postgreConfigViper.SetConfigType("yaml")
	err := postgreConfigViper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return postgreConfigInst, nil
}
