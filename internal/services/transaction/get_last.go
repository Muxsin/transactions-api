package transaction

import "infotecs-transactions-api/internal/models"

func (s *service) GetLast(count int) ([]models.Transaction, error) {
	return s.transactionRepository.GetLastByCount(count)
}
