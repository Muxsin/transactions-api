package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type transactionResponse struct {
	Id        uint   `json:"id"`
	From      string `json:"from"`
	To        string `json:"to"`
	Amount    string `json:"balance"`
	CreatedAt string `json:"created_at"`
}

func (h *handler) GetLast(c *gin.Context) {
	count, err := strconv.Atoi(c.DefaultQuery("count", "10"))
	if err != nil || count <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'count' parameter. Must be a positive integer."})
		return
	}

	transactions, err := h.getLastUseCase.Execute(count)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	var response []transactionResponse

	for _, transaction := range transactions {
		response = append(response, transactionResponse{
			Id:        transaction.ID,
			From:      transaction.From,
			To:        transaction.To,
			Amount:    fmt.Sprintf("%.2f", float64(transaction.Amount)/100),
			CreatedAt: transaction.CreatedAt.Format(time.RFC3339),
		})
	}

	c.JSON(http.StatusOK, response)
}
