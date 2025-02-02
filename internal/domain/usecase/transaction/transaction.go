package transactionusecase

import (
	"context"
	"github.com/paulparfe/finances/internal/domain/entity"
)

type Service interface {
	History(ctx context.Context, userID int) ([]entity.Transaction, error)
}

type transactionUseCase struct {
	transactionService Service
}

func NewTransactionUseCase(transactionService Service) *transactionUseCase {
	return &transactionUseCase{transactionService: transactionService}
}

func (u transactionUseCase) History(ctx context.Context, userID int) ([]entity.Transaction, error) {
	return u.transactionService.History(ctx, userID)
}
