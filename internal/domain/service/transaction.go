package service

import (
	"context"
	"errors"
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
	if userID <= 0 {
		return nil, errors.New("user_id should be greater than zero")
	}

	return s.storage.History(userID)
}
