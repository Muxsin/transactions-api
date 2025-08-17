package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	From   string
	To     string
	Amount int64
}
