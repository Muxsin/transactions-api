package send

import "infotecs-transactions-api/internal/models"

type transactionService interface {
	Send(transaction *models.Transaction) string
}

type UseCase struct {
	transactionService transactionService
}

func New(service transactionService) *UseCase {
	return &UseCase{
		transactionService: service,
	}
}
