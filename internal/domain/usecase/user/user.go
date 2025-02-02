package userusecase

import (
	"context"
	"github.com/paulparfe/finances/internal/domain/entity"
)

type Service interface {
	Deposit(ctx context.Context, dto DepositDTO) (*entity.User, error)
	Transfer(ctx context.Context, dto TransferDTO) (*entity.User, error)
}

type userUseCase struct {
	userService Service
}

func NewUserUseCase(userService Service) *userUseCase {
	return &userUseCase{userService: userService}
}

func (u userUseCase) Deposit(ctx context.Context, dto DepositDTO) (*entity.User, error) {
	return u.userService.Deposit(ctx, dto)
}

func (u userUseCase) Transfer(ctx context.Context, dto TransferDTO) (*entity.User, error) {
	return u.userService.Transfer(ctx, dto)
}
