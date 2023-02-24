package productrepository

import (
	"context"

	"github.com/ydoro/clean-go/core/domain"
	"github.com/ydoro/clean-go/core/dto"
)

func (repository repository) Create(productRequest *dto.CreateProductRequest) (*domain.Product, error) {
	ctx := context.Background()
	p := domain.Product{}
	err := repository.db.QueryRow(
		ctx,
		"INSERT INTO product (name,price,description) VALUES ($1, $2, $3) returning *",
		productRequest.Name,
		productRequest.Price,
		productRequest.Description,
	).Scan(
		&p.ID,
		&p.Name,
		&p.Price,
		&p.Description,
	)

	if err != nil {
		return nil, err
	}

	return &p, nil
}
