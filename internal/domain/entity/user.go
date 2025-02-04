package entity

import (
	"github.com/shopspring/decimal"
)

type User struct {
	ID      int             `json:"id"`
	Name    string          `json:"name"`
	Balance decimal.Decimal `json:"balance"`
}
