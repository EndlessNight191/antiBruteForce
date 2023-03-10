package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func InitRoutes() {
	e := echo.New()

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "ok")
	})

	antiBruteForce := e.Group("/api/antiBruteForce")

	AntiBrouteForceRoutes(antiBruteForce)

	e.Logger.Fatal(e.Start("localhost:" + viper.GetString("PORT")))
}
