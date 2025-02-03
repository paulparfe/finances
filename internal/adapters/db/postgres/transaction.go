package postgres

import (
	"github.com/paulparfe/finances/internal/domain/entity"
	"github.com/paulparfe/finances/pkg/client/postgresql"
)

type transactionStorage struct {
	client postgresql.Client
}

func NewTransactionStorage(client postgresql.Client) transactionStorage {
	return transactionStorage{
		client: client,
	}
}

func (r transactionStorage) History(userID int) ([]entity.Transaction, error) {
	return nil, nil
}
