package configs

import (
    "github.com/spf13/viper"
)

func init() {
    viper.SetConfigFile("../.env") // Устанавливаем имя файла конфигурации
    viper.SetConfigType("env")  // Устанавливаем тип файла конфигурации

    err := viper.ReadInConfig() // Читаем конфигурационный файл
    if err != nil {
        panic(err)
    }
}