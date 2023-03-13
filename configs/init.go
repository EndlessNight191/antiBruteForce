package configs

import (
    "github.com/spf13/viper"
)

func Init() {
    viper.SetConfigFile("../.env")
    viper.SetConfigType("env")

    if err := viper.ReadInConfig(); err != nil {
        panic(err)
    }
}