package routes

import (
	"net/http"
	"test/internal/usecase"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func useCaseMiddleware(useCase *usecase.UseCase) echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            c.Set("useCase", useCase)
            return next(c)
        }
    }
}

func InitRoutes(useCase *usecase.UseCase) {
	e := echo.New()

	e.Use(useCaseMiddleware(useCase))

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "ok")
	})

	antiBruteForce := e.Group("/api/antiBruteForce")
	admin := e.Group("/api/admin")

	AntiBrouteForceRoutes(antiBruteForce)
	AdminRoutes(admin)

	e.Logger.Fatal(e.Start(":" + viper.GetString("PORT")))
}
