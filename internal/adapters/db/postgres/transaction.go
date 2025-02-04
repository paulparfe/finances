package postgres

import (
	"context"
	"github.com/paulparfe/finances/internal/domain/entity"
	"github.com/paulparfe/finances/pkg/client/postgresql"
	"time"
)

type transactionStorage struct {
	client postgresql.Client
}

func NewTransactionStorage(client postgresql.Client) transactionStorage {
	return transactionStorage{
		client: client,
	}
}

func (s transactionStorage) History(userID int) ([]entity.Transaction, error) {
	query := `
        SELECT id, user_id, recipient_id, amount, transaction_type, created_at
        FROM transactions
        WHERE user_id = $1
        ORDER BY id DESC
        LIMIT 10`

	var transactions []entity.Transaction

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := s.client.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var transaction entity.Transaction

		err = rows.Scan(
			&transaction.ID,
			&transaction.UserID,
			&transaction.RecipientID,
			&transaction.Amount,
			&transaction.TransactionType,
			&transaction.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		transactions = append(transactions, transaction)
	}

	return transactions, nil
}
