package app

import (
	"context"
	"fmt"
	"infotecs-transactions-api/internal/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type app struct {
	config *config.Config
	router *gin.Engine
	db     *gorm.DB
	server *http.Server
}

func New(config *config.Config, db *gorm.DB) *app {
	app := &app{
		config: config,
		db:     db,
	}
	app.router = app.loadRoutes()

	app.server = &http.Server{
		Addr:    fmt.Sprintf(":%s", app.config.HTTPServerPort),
		Handler: app.router,
	}

	return app
}

func (a *app) Shutdown(ctx context.Context) error {
	if err := a.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("server shutdown failed: %w", err)
	}

	db, err := a.db.DB()
	if err != nil {
		return fmt.Errorf("failed to get database instance: %w", err)
	}

	if err := db.Close(); err != nil {
		return fmt.Errorf("failed to close database connection: %w", err)
	}

	return nil
}
