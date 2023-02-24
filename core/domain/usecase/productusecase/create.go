package productusecase

import (
	"github.com/ydoro/clean-go/core/domain"
	"github.com/ydoro/clean-go/core/dto"
)

func (uc usecase) Create(productRequest *dto.CreateProductRequest) (*domain.Product, error) {
	product, err := uc.repository.Create(productRequest)
	if err != nil {
		return nil, err
	}
	return product, nil
}
