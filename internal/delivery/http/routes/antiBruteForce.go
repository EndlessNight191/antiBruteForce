// @SwaggoFile some_service_docs2
package routes

import (
	handler "test/internal/delivery/http/handlers"
	"test/internal/usecase"

	"github.com/labstack/echo/v4"
)

func AntiBrouteForceRoutes(e *echo.Group, uc usecase.UseCase) {
	h := handler.NewAntiBrouteForceHandler(uc)

	e.POST("/antiBruteForce/", h.AntiBrouteForceCheck)
}
