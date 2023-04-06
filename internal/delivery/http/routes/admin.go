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
	l.POST("/", sHandler.AddIpToTheList)
	l.DELETE("/", sHandler.RemoveIpFromTheList)

	s := e.Group("/settings")
	s.GET("/", sHandler.GetSettings)
	s.PUT("/", sHandler.UpdateSetting)
	e.PUT("/bucket", bHandler.ResetBucket)
}
