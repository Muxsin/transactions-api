package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) Send(c *gin.Context) {
	c.JSON(http.StatusOK, h.sendUseCase.Execute())
}
