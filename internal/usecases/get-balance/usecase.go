package get_balance

import "infotecs-transactions-api/internal/models"

type walletService interface {
	GetBalance(address string) (*models.Wallet, error)
}

type UseCase struct {
	walletService walletService
}

func New(service walletService) *UseCase {
	return &UseCase{
		walletService: service,
	}
}
