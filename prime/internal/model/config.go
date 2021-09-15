package model

import (
	"tratnik.net/prime/pkg/postgres"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Database postgres.Config `mapstructure:"database"`
}
