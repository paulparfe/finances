package postgres

import (
	"database/sql"
	"github.com/paulparfe/finances/internal/domain/entity"
)

type transactionStorage struct {
	db *sql.DB
}

func NewTransactionStorage(db *sql.DB) transactionStorage {
	return transactionStorage{
		db: db,
	}
}

func (r transactionStorage) History(userID int) ([]entity.Transaction, error) {
	return nil, nil
}
