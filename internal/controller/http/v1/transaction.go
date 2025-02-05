package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/paulparfe/finances/internal/domain/entity"
	"net/http"
	"strconv"
)

type TransactionUseCase interface {
	History(ctx context.Context, userID int) ([]entity.Transaction, error)
}

type transactionHandler struct {
	transactionUseCase TransactionUseCase
}

func NewTransactionHandler(transactionUseCase TransactionUseCase) *transactionHandler {
	return &transactionHandler{
		transactionUseCase: transactionUseCase,
	}
}

func (h *transactionHandler) Register(router *gin.Engine) {
	router.GET("/users/:user_id/transactions", h.UserHistory)
}

func (h *transactionHandler) UserHistory(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user_id"})
		return
	}

	data, err := h.transactionUseCase.History(context.Background(), userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}
