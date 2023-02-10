package mappers

import (
	"database_example/core/domain/entities"
	"database_example/infrastructure/database/orm_entities"
)

func ProductToORMProduct(product entities.Product) orm_entities.ORMProduct {
	var id uint = 0
	if product.ID != nil {
		id = uint(*product.ID)
	}

	return orm_entities.ORMProduct{
		ID:        id,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
		Name:      product.Name,
		Price:     product.Price,
	}
}

func ORMProductToProduct(ormProduct orm_entities.ORMProduct) entities.Product {
	id := uint64(ormProduct.ID)
	return entities.Product{
		ID:        &id,
		CreatedAt: ormProduct.CreatedAt,
		UpdatedAt: ormProduct.UpdatedAt,
		Name:      ormProduct.Name,
		Price:     ormProduct.Price,
	}
}

func UpdateProduct(product *entities.Product, ormProduct orm_entities.ORMProduct) {
	id := uint64(ormProduct.ID)

	product.ID = &id
	product.CreatedAt = ormProduct.CreatedAt
	product.UpdatedAt = ormProduct.UpdatedAt
}
