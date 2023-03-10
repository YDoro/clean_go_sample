package productusecase

import (
	"github.com/ydoro/clean-go/core/domain"
)

type usecase struct {
	repository domain.ProductRepository
}

func New(repository domain.ProductRepository) domain.ProductUseCase {
	return &usecase{
		repository: repository,
	}
}
