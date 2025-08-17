package wallet

import "infotecs-transactions-api/internal/models"

func (r *repository) GetByAddress(address string) (*int64, error) {
	var balance int64

	result := r.db.Model(&models.Wallet{}).Select("balance").Where("address = ?", address).First(&balance)
	if result.Error != nil {
		return nil, result.Error
	}

	return &balance, nil
}
