package get_balance

import "infotecs-transactions-api/internal/models"

func (uc *UseCase) Execute(address string) (*models.Wallet, error) {
	return uc.walletService.GetBalance(address)
}
