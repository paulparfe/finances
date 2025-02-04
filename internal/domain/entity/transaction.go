package entity

import (
	"github.com/shopspring/decimal"
	"time"
)

type Transaction struct {
	ID              int             `db:"id"`
	UserID          int             `db:"userId"`
	RecipientID     *int            `db:"recipient_id"`
	Amount          decimal.Decimal `db:"amount"`
	TransactionType string          `db:"transaction_type"`
	CreatedAt       time.Time       `db:"created_at"`
}
