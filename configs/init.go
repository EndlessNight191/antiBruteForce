package configs

import (
	"github.com/spf13/viper"
)

func InitCGF() error {
	viper.SetConfigFile("../.env")
	viper.SetConfigType("env")

	return viper.ReadInConfig()
}
