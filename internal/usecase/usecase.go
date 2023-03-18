package usecase

import "test/internal/repository"

type UseCase struct {
	repo repository.Repository
	age  int
}

func NewUseCase(repo repository.Repository) *UseCase {
	return &UseCase{repo: repo}
}
