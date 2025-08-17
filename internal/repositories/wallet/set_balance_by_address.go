package wallet

import "infotecs-transactions-api/internal/models"

func (r *repository) SetBalanceByAddress(address string, balance int64) error {
	result := r.db.Model(&models.Wallet{}).
		Where("address = ?", address).
		Update("balance", balance)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
