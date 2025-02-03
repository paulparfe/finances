package postgres

import (
	"github.com/paulparfe/finances/internal/domain/entity"
	userusecase "github.com/paulparfe/finances/internal/domain/usecase/user"
	"github.com/paulparfe/finances/pkg/client/postgresql"
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
	return nil, nil
}

func (s userStorage) Transfer(dto userusecase.TransferDTO) (*entity.User, error) {
	return nil, nil
}
