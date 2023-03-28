package routes

import (
	handler "test/internal/delivery/http/handlers"
	"test/internal/usecase"

	"github.com/labstack/echo/v4"
)

func AntiBrouteForceRoutes(e *echo.Group, uc usecase.UseCase) {
	h := handler.NewAntiBrouteForceHandler(uc)

	// BRUTE FORCE
	// @Summary check brute force
	// @Description brute force check, the main feature of the application, returns true if brute force is not performed and access is allowed, or false if access is denied
	// @Tags bruteForce
	// @Accept json
	// @Produce json
	// @Param settings body domain.IncomingRequest "check brute force"
	// @Success 200 {object} Response
	// @Failure 400 {object} Response
	// @Failure 500 {object} Response
	// @Router /api/lists [post]
	e.POST("/antiBruteForce/", h.AntiBrouteForceCheck)
}
