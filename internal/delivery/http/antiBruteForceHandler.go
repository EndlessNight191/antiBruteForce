package http

import (
	"test/internal/domain"
	"test/internal/usecase"

	"github.com/labstack/echo/v4"
	"github.com/go-playground/validator/v10"
)

func AntiBrouteForceHangler(c echo.Context) error {
    req := new(domain.IncomingRequest)
    if err := c.Bind(req); err != nil {
        return err
    }

    // Validate IncomingRequest using a validator library, e.g. https://github.com/go-playground/validator
    if err := validator.New().Struct(req); err != nil {
        return err
    }

	usecase.AllowAccess(*req)
	return nil
}