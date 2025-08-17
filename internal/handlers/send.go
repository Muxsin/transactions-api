package handlers

import (
	"fmt"
	"infotecs-transactions-api/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type createTransactionRequest struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Amount string `json:"amount"`
}

func (h *handler) Send(c *gin.Context) {
	var request createTransactionRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	formattedAmount := fmt.Sprintf("%.2f", request.Amount)
	amount, err := strconv.ParseFloat(formattedAmount, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid amount format. Please provide a valid number."})
		return
	}

	transaction := &models.Transaction{
		From:   request.From,
		To:     request.To,
		Amount: int64(int(amount * 100)),
	}

	//if err := h.sendUseCase.Execute(transaction); err != nil {
	//	switch err.Error() {
	//	case "insufficient funds":
	//		c.JSON(http.StatusConflict, gin.H{"error": "insufficient funds"})
	//	case "sender wallet not found":
	//		c.JSON(http.StatusNotFound, gin.H{"error": "sender wallet not found"})
	//	default:
	//		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to complete transaction"})
	//	}
	//	return
	//}
	h.sendUseCase.Execute(transaction)

	c.Status(http.StatusCreated)
}
