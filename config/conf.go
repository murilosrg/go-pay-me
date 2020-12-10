package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

//Configuration struct
type Configuration struct {
	DB struct {
		Driver string `json:"driver"`
		Addr   string `json:"addr"`
	} `json:"db"`
	Address string `json:"address"`
}

var conf *Configuration

//Config get instance
func Config() *Configuration {
	if conf != nil {
		return conf
	}

	viper.AddConfigPath("/etc/payme/")
	viper.SetConfigName("configuration")
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatal("config file error: ", err)
	}

	if err := viper.Unmarshal(&conf); err != nil {
		logrus.Fatal("config file error: ", err)
	}

	return conf
}
