package service

import (
	"context"
	"github.com/paulparfe/finances/internal/domain/entity"
	userusecase "github.com/paulparfe/finances/internal/domain/usecase/user"
)

type UserStorage interface {
	Deposit(dto userusecase.DepositDTO) (*entity.User, error)
	Transfer(dto userusecase.TransferDTO) (*entity.User, error)
}

type userService struct {
	storage UserStorage
}

func NewUserService(storage UserStorage) *userService {
	return &userService{
		storage: storage,
	}
}

func (s userService) Deposit(ctx context.Context, dto userusecase.DepositDTO) (*entity.User, error) {
	return s.storage.Deposit(dto)
}

func (s userService) Transfer(ctx context.Context, dto userusecase.TransferDTO) (*entity.User, error) {
	return s.storage.Transfer(dto)
}
