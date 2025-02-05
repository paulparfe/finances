package service_test

import (
	"context"
	"github.com/paulparfe/finances/internal/domain/entity"
	"github.com/paulparfe/finances/internal/domain/service"
	userusecase "github.com/paulparfe/finances/internal/domain/usecase/user"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockUserStorage struct {
	mock.Mock
}

func (m *MockUserStorage) Deposit(dto userusecase.DepositDTO) (*entity.User, error) {
	args := m.Called(dto)
	return args.Get(0).(*entity.User), args.Error(1)
}

func (m *MockUserStorage) Transfer(dto userusecase.TransferDTO) (*entity.User, error) {
	args := m.Called(dto)
	return args.Get(0).(*entity.User), args.Error(1)
}

func TestDeposit_Success(t *testing.T) {
	mockStorage := new(MockUserStorage)
	userService := service.NewUserService(mockStorage)

	dto := userusecase.DepositDTO{
		UserID: 1,
		Amount: decimal.NewFromInt(100),
	}

	mockStorage.On("Deposit", dto).Return(&entity.User{
		ID:      1,
		Name:    "Adam",
		Balance: decimal.NewFromInt(100),
	}, nil)

	result, err := userService.Deposit(context.Background(), dto)

	assert.NoError(t, err)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Adam", result.Name)
	assert.Equal(t, decimal.NewFromInt(100), result.Balance)

	mockStorage.AssertExpectations(t)
}

func TestDeposit_InvalidAmount(t *testing.T) {
	mockStorage := new(MockUserStorage)
	userService := service.NewUserService(mockStorage)

	dto := userusecase.DepositDTO{
		UserID: 1,
		Amount: decimal.NewFromInt(-100),
	}

	result, err := userService.Deposit(context.Background(), dto)

	assert.Equal(t, "deposit amount must be greater than zero", err.Error())
	assert.Nil(t, result)
}

func TestTransfer_Success(t *testing.T) {
	mockStorage := new(MockUserStorage)
	userService := service.NewUserService(mockStorage)

	dto := userusecase.TransferDTO{
		SenderUserID:    1,
		RecipientUserID: 2,
		Amount:          decimal.NewFromInt(50),
	}

	mockStorage.On("Transfer", dto).Return(&entity.User{
		ID:      1,
		Name:    "Adam",
		Balance: decimal.NewFromInt(50),
	}, nil)

	result, err := userService.Transfer(context.Background(), dto)

	assert.NoError(t, err)
	assert.NotNil(t, result)

	mockStorage.AssertExpectations(t)
}

func TestTransfer_InvalidAmount(t *testing.T) {
	mockStorage := new(MockUserStorage)
	userService := service.NewUserService(mockStorage)

	dto := userusecase.TransferDTO{
		SenderUserID:    1,
		RecipientUserID: 2,
		Amount:          decimal.NewFromInt(-50),
	}

	result, err := userService.Transfer(context.Background(), dto)

	assert.Equal(t, "transfer amount must be greater than zero", err.Error())
	assert.Nil(t, result)
}

func TestTransfer_InvalidSenderUserID(t *testing.T) {
	mockStorage := new(MockUserStorage)
	userService := service.NewUserService(mockStorage)

	dto := userusecase.TransferDTO{
		SenderUserID:    0,
		RecipientUserID: 2,
		Amount:          decimal.NewFromInt(50),
	}

	result, err := userService.Transfer(context.Background(), dto)

	assert.Equal(t, "invalid sender user ID", err.Error())
	assert.Nil(t, result)
}

func TestTransfer_InvalidRecipientUserID(t *testing.T) {
	mockStorage := new(MockUserStorage)
	userService := service.NewUserService(mockStorage)

	dto := userusecase.TransferDTO{
		SenderUserID:    1,
		RecipientUserID: 0,
		Amount:          decimal.NewFromInt(50),
	}

	result, err := userService.Transfer(context.Background(), dto)

	assert.Equal(t, "invalid recipient user ID", err.Error())
	assert.Nil(t, result)
}

func TestTransfer_SenderRecipientSame(t *testing.T) {
	mockStorage := new(MockUserStorage)
	userService := service.NewUserService(mockStorage)

	dto := userusecase.TransferDTO{
		SenderUserID:    1,
		RecipientUserID: 1,
		Amount:          decimal.NewFromInt(50),
	}

	result, err := userService.Transfer(context.Background(), dto)

	assert.Equal(t, "sender and recipient user ID must be different", err.Error())
	assert.Nil(t, result)
}
