package http

import (
	"fmt"
	"net/http"
	"test/internal/domain"
	"test/internal/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func AntiBrouteForceHandler(c echo.Context) error {
    req := new(domain.IncomingRequest)
    if err := c.Bind(req); err != nil {
        return fmt.Errorf("req new struct IncomingRequest: %w", err)
    }

    if err := validator.New().Struct(req); err != nil {
        return fmt.Errorf("validator req IncomingRequest: %w", err)
    }

	access, err := usecase.AllowAccess(*req)
    if err != nil {
        return fmt.Errorf("usecase AllowAccess: %w", err)
    }

    if !access.IsAccess {
        return c.JSON(http.StatusForbidden, access)
    }
    
	return c.JSON(http.StatusOK, access)
}