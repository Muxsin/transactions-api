package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type walletResponse struct {
	Id        uint   `json:"id"`
	Address   string `json:"address"`
	Balance   string `json:"balance"`
	CreatedAt string `json:"created_at"`
}

func (h *handler) GetBalance(c *gin.Context) {
	address := c.Param("address")

	wallet, err := h.getBalanceUseCase.Execute(address)
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			c.Status(http.StatusNotFound)
			return
		}

		c.Status(http.StatusInternalServerError)
		return
	}

	response := walletResponse{
		Id:        wallet.ID,
		Address:   wallet.Address,
		Balance:   fmt.Sprintf("%.2f", float64(wallet.Balance)/100),
		CreatedAt: wallet.CreatedAt.Format(time.RFC3339),
	}

	c.JSON(http.StatusOK, response)
}
