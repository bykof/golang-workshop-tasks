package orm_entities

import (
	"time"

	"github.com/shopspring/decimal"
)

type ORMProduct struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Price     decimal.Decimal
}
