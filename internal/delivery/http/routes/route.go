package routes

import (
	"net/http"
	middle "test/internal/delivery/http/routes/middleware"
	"test/internal/domain"
	"test/internal/usecase"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"github.com/swaggo/echo-swagger"
)

// @title Anti Brute Force Service
// @description The service is designed to detect and prevent brute-force attacks on web resources.
// @version 1.0
// @host localhost:5080
// @BasePath /
func InitRoutes(useCase usecase.UseCase, config *domain.ConfigSetting, secretKey string) {
	e := echo.New()
	middleWare := middle.NewmiddleWare(secretKey)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Ping server
	// @Summary ping server
	// @Router /ping [get]
	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "ok")
	})

	antiBruteForce := e.Group("/api/antiBruteForce")
	admin := e.Group("/api/admin")
	admin.Use(middleWare.CheckBearerToken)

	AntiBrouteForceRoutes(antiBruteForce, useCase)
	AdminRoutes(admin, useCase)

	e.Logger.Fatal(e.Start(":" + viper.GetString("PORT")))
}
