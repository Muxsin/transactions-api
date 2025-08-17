package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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

	c.JSON(http.StatusOK, transactions)
}
