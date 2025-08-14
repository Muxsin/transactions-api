package app

import (
	"infotecs-transactions-api/internal/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type app struct {
	config *config.Config
	router *gin.Engine
	db     *gorm.DB
}

func New(config *config.Config, db *gorm.DB) *app {
	app := &app{
		config: config,
		db:     db,
	}
	app.router = app.loadRoutes()

	return app
}
