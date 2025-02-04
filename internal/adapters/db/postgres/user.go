package postgres

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/paulparfe/finances/internal/domain/entity"
	userusecase "github.com/paulparfe/finances/internal/domain/usecase/user"
	"github.com/paulparfe/finances/pkg/client/postgresql"
	"time"
)

type userStorage struct {
	client postgresql.Client
}

func NewUserStorage(client postgresql.Client) userStorage {
	return userStorage{
		client: client,
	}
}

func (s userStorage) Deposit(dto userusecase.DepositDTO) (*entity.User, error) {
	user := &entity.User{}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	tx, err := s.client.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.RepeatableRead})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	query := `
		INSERT INTO transactions (user_id, recipient_id, amount, transaction_type, created_at)
		VALUES ($1, NULL, $2, 'deposit', $3);
		`
	args := []any{dto.UserID, dto.Amount, time.Now()}
	_, err = tx.Exec(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	query = `
		UPDATE users
		SET balance = balance + $1
		WHERE id = $2
		RETURNING id, name, balance;
		`
	args = []any{dto.Amount, dto.UserID}
	err = tx.QueryRow(ctx, query, args...).Scan(&user.ID, &user.Name, &user.Balance)
	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "40001" {
			return nil, errors.New("transaction failed due to serialization conflict, please retry")
		}
		return nil, err
	}

	return user, nil
}

func (s userStorage) Transfer(dto userusecase.TransferDTO) (*entity.User, error) {
	return nil, nil
}
