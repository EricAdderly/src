package v1

import (
	"e-project/internal/usecase"
	models "e-project/model"
	"e-project/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReceivingTransaction struct {
	receivingTransaction usecase.ReceivingTransaction
	logger               logger.Interface
}

func newAuthRoutes(handler *gin.RouterGroup, t usecase.ReceivingTransaction, l logger.Interface) {
	r := &ReceivingTransaction{t, l}
	h := handler
	{
		h.POST("/transaction", r.doReceivingTransaction)
	}
}

func (r *ReceivingTransaction) doReceivingTransaction(c *gin.Context) {
	var request models.ReceivingTransactionRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		r.logger.Error(err, "http - v1 - doReceivingTransaction")
		errorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}

	if request.ID == "" {
		r.logger.Error("Login is empty", "http - v1 - doReceivingTransaction")
		errorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}

	if request.Sum <= 0 {
		r.logger.Error("Password is empty", "http - v1 - doReceivingTransaction")
		errorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}

	err := r.ReceivingTransaction.ReceivingTransaction(ctx, request)

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
	})

}
