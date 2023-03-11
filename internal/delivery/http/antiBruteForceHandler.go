package http

import (
	"test/internal/usecase"

	"github.com/labstack/echo/v4"
)

func AntiBrouteForceHangler(c echo.Context) error {

	usecase.AllowAccess()
	return nil
}