package ports

import (
	"product_service/core/domain"
	"product_service/core/value_objects"
)

type ProductRepositoryPort interface {
	List() (domain.Products, error)
	Get(id int64) (domain.Product, error)
	Add(product domain.Product) (int64, error)
	Update(updateProductParams value_objects.UpdateProductParams) (domain.Product, error)
	Delete(id int64) error
}
