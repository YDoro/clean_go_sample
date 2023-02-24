package domain

import (
	"net/http"

	"github.com/ydoro/clean-go/core/dto"
)

// Product is the entity of the table product in database
type Product struct {
	ID          int32
	Name        string
	Price       float32
	Description string
}

// ProductService is a contract of http adapter layer
type ProductService interface {
	Create(http.ResponseWriter, *http.Request)
	Fetch(http.ResponseWriter, *http.Request)
}

// ProductUseCase is a contract of business layer
type ProductUseCase interface {
	Create(*dto.CreateProductRequest) (*Product, error)
	Fetch(*dto.PaginationRequestParams) (*Pagination[[]Product], error)
}

// ProductRepository is a contract of databse connection adapter (infra layer)
type ProductRepository interface {
	Create(*dto.CreateProductRequest) (*Product, error)
	Fetch(*dto.PaginationRequestParams) (*Pagination[[]Product], error)
}
