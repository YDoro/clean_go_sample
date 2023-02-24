package di

import (
	"github.com/ydoro/clean-go/adapter/http/productservice"
	"github.com/ydoro/clean-go/adapter/postgres"
	"github.com/ydoro/clean-go/adapter/postgres/productrepository"
	"github.com/ydoro/clean-go/core/domain"
	"github.com/ydoro/clean-go/core/domain/usecase/productusecase"
)

func ConfigProductDI(conn postgres.PoolInterface) domain.ProductService {
	pr := productrepository.New(conn)
	puc := productusecase.New(pr)
	ps := productservice.New(puc)

	return ps
}
