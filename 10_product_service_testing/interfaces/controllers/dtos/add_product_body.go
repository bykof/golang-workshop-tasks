package dtos

import (
	"database_example/core/domain/entities"

	"github.com/shopspring/decimal"
)

type AddProductBody struct {
	Name  string  `form:"name" json:"name" xml:"name"  binding:"required"`
	Price float64 `form:"price" json:"price" xml:"price"  binding:"required"`
}

func (a AddProductBody) ToProduct() entities.Product {
	return entities.Product{
		Name:  a.Name,
		Price: decimal.NewFromFloat(a.Price),
	}
}
