package transaction

import (
	"infotecs-transactions-api/internal/models"
)

func (r *repository) Insert(transaction *models.Transaction) error {
	return r.db.Create(transaction).Error
}
