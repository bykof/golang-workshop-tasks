package entities

import (
	"github.com/shopspring/decimal"
	"time"
)

type Product struct {
	ID        *uint64
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Price     decimal.Decimal
}
