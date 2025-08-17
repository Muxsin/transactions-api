package get_last

import "infotecs-transactions-api/internal/models"

type transactionService interface {
	GetLast(count int) ([]models.Transaction, error)
}

type UseCase struct {
	transactionService transactionService
}

func New(service transactionService) *UseCase {
	return &UseCase{
		transactionService: service,
	}
}
