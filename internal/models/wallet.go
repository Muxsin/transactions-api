package models

import "gorm.io/gorm"

type Wallet struct {
	gorm.Model
	Address string
	Balance int64
}
