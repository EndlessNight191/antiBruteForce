package routes

import (
	handler "test/internal/delivery/http/handlers"

	"github.com/labstack/echo/v4"
)

func AdminRoutes(e *echo.Group) {
	l := e.Group("/lists")
	l.POST("/", handler.AddIpToTheList)
	l.DELETE("/", handler.RemoveIpFromTheList)

	s := e.Group("/settings")
	s.GET("/", handler.GetSettings)
	s.PUT("/", handler.UpdateSetting)

	e.PUT("/bucket", handler.ResetBucket)
}
