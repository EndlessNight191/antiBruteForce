package routes

import (
	handler "test/internal/delivery/http/handlers"

	"github.com/labstack/echo/v4"
)

func AdminRoutes(e *echo.Group) {
	e.POST("/lists", handler.AddIpToTheList)
	e.DELETE("/lists", handler.RemoveIpFromTheList)

	e.GET("/settings", handler.GetSettings)
	e.PUT("/settings", handler.UpdateSetting)

	e.PUT("/bucket", handler.ResetBucket)
}
