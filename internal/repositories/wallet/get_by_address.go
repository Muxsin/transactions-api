package wallet

import "infotecs-transactions-api/internal/models"

func (r *repository) GetByAddress(address string) (*models.Wallet, error) {
	var wallet models.Wallet

	result := r.db.Where("address = ?", address).First(&wallet)
	if result.Error != nil {
		return nil, result.Error
	}

	return &wallet, nil
}
