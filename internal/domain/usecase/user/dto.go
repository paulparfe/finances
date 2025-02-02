package userusecase

type DepositDTO struct {
	UserID int
	Amount int
}

type TransferDTO struct {
	SenderUserID    int
	RecipientUserID int
	Amount          int
}
