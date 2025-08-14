package app

import (
	"context"
	"fmt"
	"infotecs-transactions-api/internal/config"
	"infotecs-transactions-api/internal/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type app struct {
	config *config.Config
	router *gin.Engine
	db     *gorm.DB
	server *http.Server
}

func New() (*app, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, err
	}

	db, err := database.Connect()
	if err != nil {
		return nil, err
	}

	// region: loading routes
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Done")
	})
	// endregion

	app := &app{
		config: config.New(),
		db:     db,
		router: router,
	}

	app.server = &http.Server{
		Addr:    fmt.Sprintf(":%s", app.config.HTTPServerPort),
		Handler: app.router,
	}

	return app, nil
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
