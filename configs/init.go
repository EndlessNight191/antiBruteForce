package configs

import (
	"github.com/spf13/viper"
)

func Init() error { // InitCFG
	viper.SetConfigFile("../.env")
	viper.SetConfigType("env")

	return viper.ReadInConfig()
}
