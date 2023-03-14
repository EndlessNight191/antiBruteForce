package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

func Init() error {
    viper.SetConfigFile("../.env")
    viper.SetConfigType("env")

    if err := viper.ReadInConfig(); err != nil {
        return fmt.Errorf("error init config: %w", err)
    }
    return nil
}