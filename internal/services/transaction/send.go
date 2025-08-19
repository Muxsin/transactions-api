package transaction

import (
	"infotecs-transactions-api/internal/models"
)

func (s *service) Send(transaction *models.Transaction) error {
	return s.transactionRepository.Insert(transaction)
}
