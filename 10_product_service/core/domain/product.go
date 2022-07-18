package domain

import "github.com/shopspring/decimal"

type Product struct {
	ID       int64           `json:"id"`
	Name     string          `json:"name"`
	Category string          `json:"category"`
	Price    decimal.Decimal `json:"price"`
}

type Products []Product
