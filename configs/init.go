package configs

import (
	"test/internal/domain"

	"github.com/spf13/viper"
)

func InitCGF() (*domain.ConfigSetting, error) {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")


	return &domain.ConfigSetting{}, viper.ReadInConfig()
}
