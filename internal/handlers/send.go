package handlers

import (
	"errors"
	"infotecs-transactions-api/internal/models"
	"infotecs-transactions-api/internal/usecases/send"
	"log"
	"math"
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

	amount, err := strconv.ParseFloat(request.Amount, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid amount format. Please provide a valid number."})
		return
	}

	roundedAmount := roundToTwoDecimals(amount)

	transaction := &models.Transaction{
		From:   request.From,
		To:     request.To,
		Amount: int64(roundedAmount * 100),
	}

	if err := h.sendUseCase.Execute(transaction); err != nil {
		log.Println(err)

		if errors.Is(err, send.NotEnoughBalance) {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		} else if errors.Is(err, send.SenderWalletNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else if errors.Is(err, send.ReceiverWalletNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.Status(http.StatusInternalServerError)
		}

		return
	}

	c.Status(http.StatusCreated)
}

func roundToTwoDecimals(f float64) float64 {
	return math.Round(f*100) / 100
}
