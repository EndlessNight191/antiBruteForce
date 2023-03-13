package internal

import (
	"test/configs"
	"test/internal/routes"
	"test/internal/Infrastructure/cache"
)

func Run() {
	configs.Init()
	cache.Init()
	routes.InitRoutes()
}