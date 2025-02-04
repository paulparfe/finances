package userusecase

import "github.com/shopspring/decimal"

type DepositDTO struct {
	UserID int
	Amount decimal.Decimal
}

type TransferDTO struct {
	SenderUserID    int
	RecipientUserID int
	Amount          decimal.Decimal
}
