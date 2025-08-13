package app

import (
	"infotecs-transactions-api/internal/config"

	"github.com/gin-gonic/gin"
)

type app struct {
	config *config.Config
	router *gin.Engine
}

func New(config *config.Config) *app {
	app := &app{config: config}
	app.router = app.loadRoutes()

	return app
}
