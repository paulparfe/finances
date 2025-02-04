package dto

import (
	"github.com/shopspring/decimal"
)

type DepositDTO struct {
	Amount decimal.Decimal `json:"amount"`
}

type TransferDTO struct {
	RecipientUserID int             `json:"recipient_user_id"`
	Amount          decimal.Decimal `json:"amount"`
}
