package handlers

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type walletBalanceResponse struct {
	Address string `json:"address"`
	Balance string `json:"balance"`
}

func (h *handler) GetBalance(c *gin.Context) {
	address := c.Param("address")

	balance, err := h.getBalanceUseCase.Execute(address)
	if err != nil {
		log.Println(err)

		if errors.Is(gorm.ErrRecordNotFound, err) {
			c.Status(http.StatusNotFound)
			return
		}

		c.Status(http.StatusInternalServerError)
		return
	}

	response := walletBalanceResponse{
		Address: address,
		Balance: fmt.Sprintf("%.2f", float64(*balance)/100),
	}

	c.JSON(http.StatusOK, response)
}
