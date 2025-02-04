package entity

import (
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

type Transaction struct {
	ID              int            `db:"id"`
	UserID          int            `db:"userId"`
	RecipientID     *int           `db:"recipient_id"`
	Amount          pgtype.Numeric `db:"amount"`
	TransactionType string         `db:"transaction_type"`
	CreatedAt       time.Time      `db:"created_at"`
}
