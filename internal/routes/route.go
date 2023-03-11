package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func InitRoutes() {
	e := echo.New()

	antiBruteForce := e.Group("/api/antiBruteForce")

	AntiBrouteForceRoutes(antiBruteForce)

	e.Logger.Fatal(e.Start(":" + viper.GetString("PORT")))
}