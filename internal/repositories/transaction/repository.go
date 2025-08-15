package transaction

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *repository {
	return &repository{db: db}
}
