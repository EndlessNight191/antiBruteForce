package internal

import (
	"log"
	"test/configs"
	"test/internal/delivery/http/routes"
	"test/internal/repository"
	"test/internal/usecase"
)

func Run() {
	if err := configs.InitCGF(); err != nil {
		log.Fatal(err)
	}

	redisClient, err := repository.InitCache()
	if err != nil {
		log.Fatalf("error trying init cache: %v", err)
	}


	repo := repository.NewRepository(redisClient)
	useCase := usecase.NewUseCase(repo)

	routes.InitRoutes(useCase)
}
