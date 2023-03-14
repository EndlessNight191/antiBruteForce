package internal

import (
	"log"
	"test/configs"
	"test/internal/Infrastructure/cache"
	"test/internal/routes"
)

func Run() {
	if err := configs.Init(); err != nil {
		log.Fatal(err)
	}
	if err := cache.Init(); err != nil {
		log.Fatal(err)
	}
	routes.InitRoutes()
}