package wallet

import "infotecs-transactions-api/internal/models"

type walletRepository interface {
	GetByAddress(address string) (*models.Wallet, error)
}

type service struct {
	walletRepository walletRepository
}

func New(repository walletRepository) *service {
	return &service{
		walletRepository: repository,
	}
}
