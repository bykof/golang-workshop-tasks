package dtos

import (
	"database_example/core/domain/entities"

	"github.com/shopspring/decimal"
)

type UpdateProductBody struct {
	Name  string  `form:"name" json:"name" xml:"name"  binding:"required"`
	Price float64 `form:"price" json:"price" xml:"price"  binding:"required"`
}

func (u UpdateProductBody) ToProduct(id uint64) entities.Product {
	return entities.Product{
		ID:    &id,
		Name:  u.Name,
		Price: decimal.NewFromFloat(u.Price),
	}
}
