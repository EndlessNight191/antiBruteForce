package routes

import (
	handler "test/internal/delivery/http/handlers"

	"github.com/labstack/echo/v4"
)

func AntiBrouteForceRoutes(e *echo.Group) {
	e.POST("/", handler.AntiBrouteForceHandler)
}
