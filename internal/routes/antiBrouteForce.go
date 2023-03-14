package routes

import (
	"test/internal/delivery/http"

	"github.com/labstack/echo/v4"
)

func AntiBrouteForceRoutes(e *echo.Group) {
	e.POST("/", http.AntiBrouteForceHandler)
}