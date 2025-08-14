package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetLast(c *gin.Context) {
	c.JSON(http.StatusOK, h.getLastUseCase.Execute())
}
