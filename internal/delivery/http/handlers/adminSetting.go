package http

import (
	"net/http"
	"test/internal/domain"
	"test/internal/usecase"

	"github.com/labstack/echo/v4"
)

type settingsHandler struct {
	usecase usecase.UseCase
}

func NewSettingsHandler(uc usecase.UseCase) *settingsHandler {
	return &settingsHandler{
		usecase: uc,
	}
}

func (uc settingsHandler) GetSettings(c echo.Context) error {
	settings, err := uc.usecase.GetSettings()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{message: "internal server error"})
	}

	return c.JSON(http.StatusOK, settings)
}

func (uc settingsHandler) UpdateSetting(c echo.Context) error {
    updateSetting := new(domain.ConfigSetting)
    if err := c.Bind(&updateSetting); err != nil {
        return c.JSON(http.StatusBadRequest, Response{message: "validator req ConfigSetting"})
    }

	if err := uc.usecase.UpdateSettings(*updateSetting); err != nil {
		return c.JSON(http.StatusInternalServerError, Response{message: "internal server error"})
	}
	return c.JSON(http.StatusOK, updateSetting)
}

func (uc settingsHandler) AddIpToTheList(c echo.Context) error {
	listsActions := new(domain.ListsActions)
    if err := c.Bind(&listsActions); err != nil {
        return c.JSON(http.StatusBadRequest, Response{message: "validator req listsActions"})
    }

	if err := uc.usecase.AddIpToTheList(*listsActions); err != nil {
		return c.JSON(http.StatusBadRequest, Response{message: "internal server error"})
	}
	return c.JSON(http.StatusOK, Response{message: "ok"})
}

func (uc settingsHandler) RemoveIpFromTheList(c echo.Context) error {
	listsActions := new(domain.ListsActions)
    if err := c.Bind(&listsActions); err != nil {
        return c.JSON(http.StatusBadRequest, Response{message: "validator req listsActions"})
    }

	if err := uc.usecase.RemoveIpToTheList(*listsActions); err != nil {
		return c.JSON(http.StatusBadRequest, Response{message: "internal server error"})
	}
	return c.JSON(http.StatusOK, Response{message: "ok"})
}