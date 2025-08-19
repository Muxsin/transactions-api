package services

import "gorm.io/gorm"

type dbTransaction struct {
	db *gorm.DB
	tx *gorm.DB
}

func New(db *gorm.DB) *dbTransaction {
	return &dbTransaction{db: db}
}

func (t *dbTransaction) Begin() {
	t.tx = t.db.Begin()
}

func (t *dbTransaction) Rollback() {
	t.tx.Rollback()
}

func (t *dbTransaction) Commit() error {
	return t.tx.Commit().Error
}
