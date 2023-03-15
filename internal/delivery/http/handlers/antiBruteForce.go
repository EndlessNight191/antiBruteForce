package http

import (
	"net/http"
	"test/internal/domain"
	"test/internal/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func AntiBrouteForce(c echo.Context) error {
    req := new(domain.IncomingRequest)
    if err := c.Bind(&req); err != nil {
        return c.JSON(http.StatusBadRequest, "validator req IncomingRequest")
    }

    if err := validator.New().Struct(&req); err != nil {
        return c.JSON(http.StatusBadRequest, "validator req IncomingRequest")
    }

	access, err := usecase.AllowAccess(*req)
    if err != nil {
        log.Info("usecase AllowAccess: %w", err)
        return c.JSON(http.StatusInternalServerError, "server error")
    }

    if !access.IsAccess {
        return c.JSON(http.StatusForbidden, access)
    }
    
	return c.JSON(http.StatusOK, access)
}