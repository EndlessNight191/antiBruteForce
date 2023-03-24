package usecase

import (
	"test/internal/domain"
	"test/internal/repository"
)

type UseCase struct {
    repo repository.Repository
    setting *domain.ConfigSetting
}

func NewUseCase(repo repository.Repository, setting *domain.ConfigSetting) *UseCase {
    return &UseCase{repo, setting}
}