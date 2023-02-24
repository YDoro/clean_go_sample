package productrepository

import (
	"github.com/ydoro/clean-go/adapter/postgres"
	"github.com/ydoro/clean-go/core/domain"
)

type repository struct {
	db postgres.PoolInterface
}

// New Returns contract implementation of ProductRepository
func New(db postgres.PoolInterface) domain.ProductRepository {
	return &repository{
		db: db,
	}
}
