package send

import "infotecs-transactions-api/internal/models"

func (uc *UseCase) Execute(transaction *models.Transaction) string {
	return uc.transactionService.Send(transaction)
}
