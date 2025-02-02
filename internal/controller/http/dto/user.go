package dto

type DepositDTO struct {
	Amount int `json:"amount"`
}

type TransferDTO struct {
	SenderUserID    int `json:"sender_user_id"`
	RecipientUserID int `json:"recipient_user_id"`
	Amount          int `json:"amount"`
}
