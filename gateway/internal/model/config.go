package model

import (
	"tratnik.net/gateway/pkg/http/server"
)

type Config struct {
	Server server.Config `mapstructure:"server"`
	Prime  struct {
		Address string `mapstructure:"address"`
	} `mapstructure:"prime"`
}
