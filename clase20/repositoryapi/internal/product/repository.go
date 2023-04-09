package product

import (
	"fmt"
	"repositoryapi/internal/domain"
	"repositoryapi/pkg/store"

	"repositoryapi/pkg/web"
)

type IRepository interface{
	GetByID(id int) (*domain.Product, error)
}

type Repository struct {
	Storage store.StoreInterface
}

func (r *Repository) GetByID(id int) (*domain.Product, error) {
	prod, err := r.Storage.Read(id)
	if err != nil {
		return nil, web.NewNotFoundApiError(fmt.Sprintf("product_id %d not found", id))
	}
	return prod, nil
}
