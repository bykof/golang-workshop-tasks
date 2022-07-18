package domain_services

import (
	"product_service/core/domain"
	"product_service/core/value_objects"

	"github.com/shopspring/decimal"
)

func ProductBodyToProduct(productBody value_objects.ProductBody) domain.Product {
	return domain.Product{
		Name:     productBody.Name,
		Category: productBody.Category,
		Price:    decimal.NewFromFloat(productBody.Price),
	}
}

func ProductBodyToUpdateProductParams(id int64, productBody value_objects.ProductBody) value_objects.UpdateProductParams {
	return value_objects.UpdateProductParams{
		ID:       id,
		Name:     productBody.Name,
		Price:    decimal.NewFromFloat(productBody.Price),
		Category: productBody.Category,
	}
}
