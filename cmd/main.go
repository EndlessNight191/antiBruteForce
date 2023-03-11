package main

import (
	"test/internal/routes"
	"test/internal/Infrastructure/cache"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func main() {

	cache.Init()
	routes.InitRoutes();
}