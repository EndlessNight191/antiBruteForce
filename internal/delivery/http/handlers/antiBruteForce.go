package http

import (
	"net/http"
	"test/internal/domain"
	"test/internal/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type Response struct {
    message string
}

func AntiBrouteForce(c echo.Context) error {
    req := new(domain.IncomingRequest)
    if err := c.Bind(&req); err != nil {
        return c.JSON(http.StatusBadRequest, Response{message: "validator req IncomingRequest"})
    }

    if err := validator.New().Struct(&req); err != nil {
        return c.JSON(http.StatusBadRequest, Response{message: "validator req IncomingRequest"})
    }

    useCaseAntiBrouteForce := c.Get("useCase").(*usecase.UseCase)

	access, err := useCaseAntiBrouteForce.AllowAccess(*req)
    if err != nil {
        log.Error("usecase AllowAccess: %w", err)
        return c.JSON(http.StatusInternalServerError,Response{message: "validator req IncomingRequest"})
    }

    if !access.IsAccess {
        return c.JSON(http.StatusOK, access)
    }
    
	return c.JSON(http.StatusOK, access)
}