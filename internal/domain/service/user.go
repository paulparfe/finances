package service

import (
	"context"
	"errors"
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
	if dto.Amount.Sign() <= 0 {
		return nil, errors.New("deposit amount must be greater than zero")
	}

	return s.storage.Deposit(dto)
}

func (s userService) Transfer(ctx context.Context, dto userusecase.TransferDTO) (*entity.User, error) {
	if dto.Amount.Sign() <= 0 {
		return nil, errors.New("transfer amount must be greater than zero")
	}

	if dto.SenderUserID <= 0 {
		return nil, errors.New("invalid sender user ID")
	}

	if dto.RecipientUserID <= 0 {
		return nil, errors.New("invalid recipient user ID")
	}

	if dto.SenderUserID == dto.RecipientUserID {
		return nil, errors.New("sender and recipient user ID must be different")
	}

	return s.storage.Transfer(dto)
}
