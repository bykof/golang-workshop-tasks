package value_objects

import "github.com/shopspring/decimal"

type UpdateProductParams struct {
	ID       int64
	Name     string
	Price    decimal.Decimal
	Category string
}
