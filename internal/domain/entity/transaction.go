package entity

type Transaction struct {
	ID     int `json:"id"`
	UserID int `json:"userId"`
	Amount int `json:"amount"`
}
