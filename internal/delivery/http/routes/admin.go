// @SwaggoFile some_service_docs
package routes

import (
	handler "test/internal/delivery/http/handlers"
	"test/internal/usecase"

	"github.com/labstack/echo/v4"
)

func AdminRoutes(e *echo.Group, uc usecase.UseCase) {
	sHandler := handler.NewSettingsHandler(uc)
	bHandler := handler.NewBacketHandler(uc)

	l := e.Group("/lists")
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
	l.POST("/", sHandler.AddIpToTheList)

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
	l.DELETE("/", sHandler.RemoveIpFromTheList)

	s := e.Group("/settings")
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
	s.GET("/", sHandler.GetSettings)

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
	s.PUT("/", sHandler.UpdateSetting)

	// ResetBucket
	// @Summary Reset the bucket
	// @Description Reset the state of the bucket to its initial state
	// @Tags adminBucket
	// @Accept json
	// @Produce json
	// @Param resetBucket body domain.ResetBucket true "The bucket to reset"
	// @Success 200 {object} Response
	// @Failure 400 {object} Response
	// @Failure 500 {object} Response
	// @Router /api/admins/bucket/ [put]
	e.PUT("/bucket", bHandler.ResetBucket)
}
