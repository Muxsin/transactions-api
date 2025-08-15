package app

import (
	"context"
	"fmt"
	"infotecs-transactions-api/internal/config"
	"infotecs-transactions-api/internal/database"
	"infotecs-transactions-api/internal/handlers"
	transaction_repository "infotecs-transactions-api/internal/repositories/transaction"
	wallet_repository "infotecs-transactions-api/internal/repositories/wallet"
	"infotecs-transactions-api/internal/services/transaction"
	"infotecs-transactions-api/internal/services/wallet"
	get_balance "infotecs-transactions-api/internal/usecases/get-balance"
	get_last "infotecs-transactions-api/internal/usecases/get-last"
	"infotecs-transactions-api/internal/usecases/send"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type app struct {
	config *config.Config
	router *gin.Engine
	// Todo: remove from struct and move in to method New() when implements repository.
	db     *gorm.DB
	server *http.Server
}

func New() (*app, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, err
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	db, err := database.Connect(dbHost, dbUser, dbPassword, dbName, dbPort)
	if err != nil {
		return nil, err
	}

	// region: creating repositories
	walletRepository := wallet_repository.New(db)
	transactionRepository := transaction_repository.New(db)
	// endregion

	// region: creating services
	walletService := wallet.New(walletRepository)
	transactionService := transaction.New(transactionRepository)
	// endregion

	// region: creating usecases
	getBalanceUseCase := get_balance.New(walletService)
	getLastUseCase := get_last.New(transactionService)
	sendUseCase := send.New(transactionService)
	// endregion

	// region: creating handlers
	handler := handlers.New(getBalanceUseCase, getLastUseCase, sendUseCase)
	// endregion

	// region: loading routes
	router := gin.Default()
	router.GET("/api/wallet/:address/balance", handler.GetBalance)
	router.GET("/api/transactions", handler.GetLast)
	router.POST("/api/send", handler.Send)
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
