package handlers

import (
	"net/http"
	"test/internal/domain"
	"test/internal/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type backetHandler struct {
	usecase usecase.UseCase
}

func NewBacketHandler(uc usecase.UseCase) *backetHandler {
	return &backetHandler{
		usecase: uc,
	}
}

func (uc backetHandler) ResetBucket(c echo.Context) error {
	resetBucket := new(domain.ResetBucket)
    if err := c.Bind(&resetBucket); err != nil {
        return c.JSON(http.StatusBadRequest, Response{Message: "validator req listsActions"})
    }

	if err := validator.New().Struct(resetBucket); err != nil {
        return c.JSON(http.StatusBadRequest, Response{Message: err.Error()})
    }

	if err := uc.usecase.ResetBucket(*resetBucket); err != nil {
		return c.JSON(http.StatusBadRequest, Response{Message: "internal server error"})
	}
	return c.JSON(http.StatusOK, Response{Message: "ok"})
}