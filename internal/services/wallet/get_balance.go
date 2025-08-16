package wallet

import "infotecs-transactions-api/internal/models"

func (s *service) GetBalance(address string) (*models.Wallet, error) {
	return s.walletRepository.GetByAddress(address)
}
