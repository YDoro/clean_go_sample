package productrepository

import (
	"context"

	"github.com/booscaaa/go-paginate/paginate"
	"github.com/ydoro/clean-go/core/domain"
	"github.com/ydoro/clean-go/core/dto"
)

func (r repository) Fetch(pagination *dto.PaginationRequestParams) (*domain.Pagination[[]domain.Product], error) {
	ctx := context.Background()
	products := []domain.Product{}
	total := int32(0)

	pagin := paginate.Instance(domain.Product{})
	query, queryCount := pagin.Query("SELECT * FROM product").
		Page(pagination.Page).
		Desc(pagination.Descending).
		Sort(pagination.Sort).
		RowsPerPage(pagination.ItemsPerPage).
		SearchBy(pagination.Search, "name", "description").
		Select()
	{
		rows, err := r.db.Query(ctx, *query)
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			p := domain.Product{}
			rows.Scan(
				&p.ID,
				&p.Name,
				&p.Price,
				&p.Description,
			)
			products = append(products, p)
		}

	}
	{
		err := r.db.QueryRow(ctx, *queryCount).Scan(&total)
		if err != nil {
			return nil, err
		}
		return &domain.Pagination[[]domain.Product]{
			Items: products,
			Total: total,
		}, nil
	}
}
