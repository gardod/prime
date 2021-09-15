package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"tratnik.net/gateway/internal/model"
)

func GetFromFile() *model.Config {
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/gateway/")
	viper.AddConfigPath("./config/")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		logrus.WithError(err).Fatal("Unable to read config")
	}

	config := &model.Config{}
	if err := viper.Unmarshal(&config); err != nil {
		logrus.WithError(err).Fatal("Unable to unmarshal config")
	}

	return config
}
