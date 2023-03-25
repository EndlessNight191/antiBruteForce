package internal

import (
	"log"
	"test/configs"
	"test/internal/delivery/http/routes"
	"test/internal/repository"
	"test/internal/usecase"
)

func Run() {
	configSetting, err := configs.InitCGF()
	if err != nil {
		log.Fatal(err)
	}


	redisClient, err := repository.InitCache()
	if err != nil {
		log.Fatalf("error trying init cache: %v", err)
	}


	repo := repository.NewRepository(redisClient, configSetting)
	useCase := usecase.NewUseCase(repo, configSetting)
	*configSetting, err = useCase.GetSettings()
	if err != nil {
		log.Fatalf("error trying init setting: %v", err)
	}

	routes.InitRoutes(useCase)
}
