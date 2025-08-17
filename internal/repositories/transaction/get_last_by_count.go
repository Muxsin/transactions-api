package transaction

import "infotecs-transactions-api/internal/models"

func (r *repository) GetLastByCount(count int) ([]models.Transaction, error) {
	var transactions []models.Transaction

	if err := r.db.Order("created_at desc").Limit(count).Find(&transactions).Error; err != nil {
		return nil, err
	}

	return transactions, nil
}
