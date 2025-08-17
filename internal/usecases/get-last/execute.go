package get_last

import "infotecs-transactions-api/internal/models"

func (uc *UseCase) Execute(count int) ([]models.Transaction, error) {
	return uc.transactionService.GetLast(count)
}
