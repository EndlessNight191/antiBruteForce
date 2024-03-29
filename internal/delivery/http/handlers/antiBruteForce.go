package handlers

import (
	"net/http"
	"test/internal/domain"
	"test/internal/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type Response struct {
    Message string `json:"message"`
}

type antiBrouteForceHandler struct {
	usecase usecase.UseCase
}

func NewAntiBrouteForceHandler(uc usecase.UseCase) *antiBrouteForceHandler {
	return &antiBrouteForceHandler{
		usecase: uc,
	}
}

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
func (uc antiBrouteForceHandler) AntiBrouteForceCheck(c echo.Context) error {
    req := new(domain.IncomingRequest)
    if err := c.Bind(&req); err != nil {
        return c.JSON(http.StatusBadRequest, Response{Message: "validator req IncomingRequest"})
    }

    if err := validator.New().Struct(req); err != nil {
        return c.JSON(http.StatusBadRequest, Response{Message: err.Error()})
    }

	access, err := uc.usecase.AllowAccess(*req)
    if err != nil {
        log.Error("usecase AllowAccess: %w", err)
        return c.JSON(http.StatusInternalServerError, Response{Message: "validator req IncomingRequest"})
    }

    if !access.IsAccess {
        return c.JSON(http.StatusOK, access)
    }
    
	return c.JSON(http.StatusOK, access)
}