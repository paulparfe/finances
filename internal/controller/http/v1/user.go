package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/paulparfe/finances/internal/controller/http/dto"
	"github.com/paulparfe/finances/internal/domain/entity"
	userusecase "github.com/paulparfe/finances/internal/domain/usecase/user"
	"net/http"
	"strconv"
)

type UserUseCase interface {
	Deposit(ctx context.Context, dto userusecase.DepositDTO) (*entity.User, error)
	Transfer(ctx context.Context, dto userusecase.TransferDTO) (*entity.User, error)
}

type userHandler struct {
	userUseCase UserUseCase
}

func NewUserHandler(userUseCase UserUseCase) *userHandler {
	return &userHandler{
		userUseCase: userUseCase,
	}
}

func (u *userHandler) Register(router *gin.Engine) {
	router.POST("/users/:user_id/deposit", u.Deposit)
	router.POST("/users/:user_id/transfer", u.Transfer)
}

func (u *userHandler) Deposit(c *gin.Context) {
	// Get userID
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id"})
		return
	}
	if userID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id should be greater than zero"})
		return
	}

	// Get amount
	controllerDTO := dto.DepositDTO{}
	if err := c.ShouldBindJSON(&controllerDTO); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	// Prepare transactionusecase.HistoryDTO
	useCaseDTO := userusecase.DepositDTO{
		UserID: userID,
		Amount: controllerDTO.Amount,
	}

	data, err := u.userUseCase.Deposit(context.Background(), useCaseDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}

func (u *userHandler) Transfer(c *gin.Context) {
	// Get userID
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id"})
		return
	}

	// Get amount and recipient user_id
	transferDTO := dto.TransferDTO{}
	if err := c.ShouldBindJSON(&transferDTO); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	// Prepare transactionusecase.TransferDTO
	useCaseDTO := userusecase.TransferDTO{
		SenderUserID:    userID,
		Amount:          transferDTO.Amount,
		RecipientUserID: transferDTO.RecipientUserID,
	}

	data, err := u.userUseCase.Transfer(context.Background(), useCaseDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}
