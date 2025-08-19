package transaction

import "infotecs-transactions-api/internal/models"

type transactionRepository interface {
	GetLastByCount(count int) ([]models.Transaction, error)
	Insert(transaction *models.Transaction) error
}

type service struct {
	transactionRepository transactionRepository
}

func New(repository transactionRepository) *service {
	return &service{
		transactionRepository: repository,
	}
}
