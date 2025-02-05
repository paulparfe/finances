package service_test

import (
	"context"
	"github.com/paulparfe/finances/internal/domain/entity"
	"github.com/paulparfe/finances/internal/domain/service"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

type MockTransactionStorage struct {
	mock.Mock
}

func (m *MockTransactionStorage) History(userID int) ([]entity.Transaction, error) {
	args := m.Called(userID)
	return args.Get(0).([]entity.Transaction), args.Error(1)
}

func TestHistory_Success(t *testing.T) {
	mockStorage := new(MockTransactionStorage)

	transactionService := service.NewTransactionService(mockStorage)

	transactions := []entity.Transaction{
		{
			ID:              1,
			UserID:          1,
			RecipientID:     nil,
			Amount:          decimal.NewFromInt(123),
			TransactionType: "deposit",
			CreatedAt:       time.Unix(0, 0),
		},
	}

	mockStorage.On("History", 1).Return(transactions, nil)

	result, err := transactionService.History(context.Background(), 1)

	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, 1, result[0].ID)
	assert.Equal(t, 1, result[0].UserID)
	assert.Nil(t, result[0].RecipientID)
	assert.Equal(t, decimal.NewFromInt(123), result[0].Amount)
	assert.Equal(t, "deposit", result[0].TransactionType)
	assert.Equal(t, time.Unix(0, 0), result[0].CreatedAt)

	mockStorage.AssertExpectations(t)
}

func TestHistory_InvalidUserID(t *testing.T) {
	mockStorage := new(MockTransactionStorage)

	transactionService := service.NewTransactionService(mockStorage)

	result, err := transactionService.History(context.Background(), 0)

	assert.Error(t, err)
	assert.Nil(t, result)
}
