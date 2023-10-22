package model

import "gorm.io/gorm"

type CurrencyTransaction struct {
	gorm.Model
	UserID           string
	Username         string
	TransactionType  string
	AvaliableBalance float64
	TransactionData  float64
	NewBalance       float64
}
