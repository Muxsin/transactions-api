package send

import (
	"infotecs-transactions-api/internal/models"
)

type transactionService interface {
	Send(transaction *models.Transaction) error
}

type walletService interface {
	GetBalance(address string) (*int64, error)
	SetBalance(address string, balance int64) error
}

type dbTransactionService interface {
	Begin()
	Rollback()
	Commit() error
}

type UseCase struct {
	transactionService   transactionService
	walletService        walletService
	dbTransactionService dbTransactionService
}

func New(transactionService transactionService, walletService walletService, dbTransactionService dbTransactionService) *UseCase {
	return &UseCase{
		transactionService:   transactionService,
		walletService:        walletService,
		dbTransactionService: dbTransactionService,
	}
}
