package routes

import (
	"net/http"
	"test/internal/usecase"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func InitRoutes(useCase *usecase.UseCase) {
	e := echo.New()

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "ok")
	})

	antiBruteForce := e.Group("/api/antiBruteForce")
	admin := e.Group("/api/admin")

	AntiBrouteForceRoutes(antiBruteForce, *useCase)
	AdminRoutes(admin)

	e.Logger.Fatal(e.Start(":" + viper.GetString("PORT")))
}
