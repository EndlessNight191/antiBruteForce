package routes

import (
	handler "test/internal/delivery/http/handlers"

	"github.com/labstack/echo/v4"
)

func AntiBrouteForceRoutes(e *echo.Group) {
	a := handler.NewSome()

	e.POST("/", a.AntiBrouteForce)
}
