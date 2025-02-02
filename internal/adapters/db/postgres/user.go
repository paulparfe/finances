package postgres

import (
	"database/sql"
	"github.com/paulparfe/finances/internal/domain/entity"
	userusecase "github.com/paulparfe/finances/internal/domain/usecase/user"
)

type userStorage struct {
	db *sql.DB
}

func NewUserStorage(db *sql.DB) userStorage {
	return userStorage{
		db: db,
	}
}

func (s userStorage) Deposit(dto userusecase.DepositDTO) (*entity.User, error) {
	return nil, nil
}

func (s userStorage) Transfer(dto userusecase.TransferDTO) (*entity.User, error) {
	return nil, nil
}
