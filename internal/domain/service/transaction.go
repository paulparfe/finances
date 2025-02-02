package service

import (
	"context"
	"github.com/paulparfe/finances/internal/domain/entity"
)

type TransactionStorage interface {
	History(userID int) ([]entity.Transaction, error)
}

type transactionService struct {
	storage TransactionStorage
}

func NewTransactionService(storage TransactionStorage) transactionService {
	return transactionService{
		storage: storage,
	}
}

func (s transactionService) History(ctx context.Context, userID int) ([]entity.Transaction, error) {
	return s.storage.History(userID)
}
