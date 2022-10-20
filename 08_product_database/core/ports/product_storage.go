package ports

import "database_example/core/domain/entities"

type ProductRepository interface {
	List() []entities.Product
	Add(product *entities.Product) error
	Save(product *entities.Product) error
	Remove(product *entities.Product) error
}
