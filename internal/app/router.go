package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *app) loadRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Done")
	})

	return router
}
