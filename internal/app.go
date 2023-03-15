package internal

import (
	"log"
	"test/configs"
	"test/internal/repository"
	"test/internal/delivery/http/routes"
)

func Run() {
	if err := configs.InitCGF(); err != nil {
		log.Fatal(err)
	}

	if err := repository.InitCache(); err != nil {
		log.Fatalf("error trying init cache: %v", err)
	}

	routes.InitRoutes()
}
