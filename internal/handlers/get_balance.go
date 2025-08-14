package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetBalance(c *gin.Context) {
	c.JSON(http.StatusOK, h.getBalanceUseCase.Execute())
}
