package middleware

import (
	"net/http"
	"strings"
	"test/internal/delivery/http/handlers"

	"github.com/labstack/echo/v4"
)

type middleWare struct {
	secretKey string
}

func NewmiddleWare(key string) *middleWare {
	return &middleWare{
		secretKey: key,
	}
}

func (key middleWare) CheckBearerToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
        token := c.Request().Header.Get("Authorization")
        if token == "" {
            return c.JSON(http.StatusUnauthorized, handlers.Response{Message: "Unauthorized"})
        }

        bearerToken := strings.Split(token, "Bearer ")
        if len(bearerToken) != 2 {
            return c.JSON(http.StatusBadRequest, handlers.Response{Message: "Invalid Authorization header format"})
        }

        if bearerToken[1] != key.secretKey {
            return c.JSON(http.StatusUnauthorized, handlers.Response{Message: "Unauthorized"})
        }

        return next(c)
    }
}
