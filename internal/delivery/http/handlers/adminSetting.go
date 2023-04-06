package handlers

import (
	"net/http"
	"test/internal/domain"
	"test/internal/usecase"

	"github.com/go-playground/validator/v10"
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

// GET SETTINGS
// @Summary Get settings in app
// @Description get the actual settings that affect the operation of the application
// @Tags adminSettings
// @Accept json
// @Produce json
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /api/admins/settings/ [get]
func (uc settingsHandler) GetSettings(c echo.Context) error {
	settings, err := uc.usecase.GetSettings()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Message: "internal server error"})
	}

	return c.JSON(http.StatusOK, settings)
}

// UPDATE SETTINGS
// @Summary update settings in app
// @Description update the actual settings that affect the operation of the application
// @Tags adminSettings
// @Accept json
// @Produce json
// @Param settings body domain.ConfigSettingtrue "update settings"
// @Success 200 domain.ConfigSettingtrue Response
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /api/admins/settings/ [put]
func (uc settingsHandler) UpdateSetting(c echo.Context) error {
    updateSetting := new(domain.ConfigSetting)
    if err := c.Bind(&updateSetting); err != nil {
        return c.JSON(http.StatusBadRequest, Response{Message: "validator req ConfigSetting"})
    }

	if err := validator.New().Struct(updateSetting); err != nil {
        return c.JSON(http.StatusBadRequest, Response{Message: err.Error()})
    }

	if err := uc.usecase.UpdateSettings(*updateSetting); err != nil {
        return c.JSON(http.StatusBadRequest, Response{Message: err.Error()})
    }
	return c.JSON(http.StatusOK, updateSetting)
}

// LISTS
// @Summary add ip in black/white lists
// @Description adds an ip address to the white/black lists for quick access or a ban on brute force
// @Tags adminLists
// @Accept json
// @Produce json
// @Param settings body domain.ListsActions "add ip in lists"
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /api/admins/lists/ [post]
func (uc settingsHandler) AddIpToTheList(c echo.Context) error {
	listsActions := new(domain.ListsActions)
    if err := c.Bind(&listsActions); err != nil {
        return c.JSON(http.StatusBadRequest, Response{Message: "validator req listsActions"})
    }

	if err := validator.New().Struct(listsActions); err != nil {
        return c.JSON(http.StatusBadRequest, Response{Message: err.Error()})
    }

	if err := uc.usecase.AddIpToTheList(*listsActions); err != nil {
		return c.JSON(http.StatusBadRequest, Response{Message: "internal server error"})
	}
	return c.JSON(http.StatusOK, Response{Message: "ok"})
}

// LISTS
// @Summary delete ip from black/white lists
// @Description removes the ip address from the white/black lists for quick access or brute force ban
// @Tags adminLists
// @Accept json
// @Produce json
// @Param settings body domain.ListsActions "delete ip from lists"
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /api/admins/lists/ [delete]
func (uc settingsHandler) RemoveIpFromTheList(c echo.Context) error {
	listsActions := new(domain.ListsActions)
    if err := c.Bind(&listsActions); err != nil {
        return c.JSON(http.StatusBadRequest, Response{Message: "validator req listsActions"})
    }

	if err := validator.New().Struct(listsActions); err != nil {
        return c.JSON(http.StatusBadRequest, Response{Message: err.Error()})
    }

	if err := uc.usecase.RemoveIpToTheList(*listsActions); err != nil {
		return c.JSON(http.StatusBadRequest, Response{Message: "internal server error"})
	}
	return c.JSON(http.StatusOK, Response{Message: "ok"})
}