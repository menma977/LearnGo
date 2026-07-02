package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Item struct {
	ID          int             `json:"-" db:"id"`
	Uuid        string          `json:"id" db:"uuid"`
	Name        string          `json:"name" db:"name"`
	Description string          `json:"description" db:"description"`
	Price       decimal.Decimal `json:"price" db:"price"`
	Quantity    int64           `json:"quantity" db:"quantity"`
	CreatedById int64           `json:"-" db:"created_by"`
	UpdatedById *int64          `json:"-" db:"updated_by"`
	DeletedById *int64          `json:"-" db:"deleted_by"`
	CreatedAt   *time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt   *time.Time      `json:"updated_at" db:"updated_at"`
	DeletedAt   *time.Time      `json:"deleted_at" db:"deleted_at"`

	CreatedBy *User `json:"created_by,omitzero" db:"-"`
	UpdatedBy *User `json:"updated_by,omitzero" db:"-"`
	DeletedBy *User `json:"deleted_by,omitzero" db:"-"`
}
