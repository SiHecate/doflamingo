package controllers

import (
	"doflamingo/database"
	"doflamingo/model"
)

func CurrencyTransaction(UserID string, TransactionType string, AvaliableBalance float64, TransactionData float64) {

	old_balance := AvaliableBalance
	new_balance := old_balance + TransactionData

	var Username string
	var user model.User
	database.Conn.Model(&user).Where("id = ?", UserID).Pluck("username", &Username)

	transaction := model.CurrencyTransaction{
		UserID:           UserID,
		Username:         Username,
		TransactionType:  TransactionType,
		AvaliableBalance: AvaliableBalance,
		TransactionData:  TransactionData,
		NewBalance:       new_balance,
	}

	database.Conn.Create(&transaction)

}
