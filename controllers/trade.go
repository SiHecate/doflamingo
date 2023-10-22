package controllers

import (
	"doflamingo/database"
	helpers "doflamingo/helper"
	"doflamingo/model"

	"github.com/gofiber/fiber/v2"
)

func AddBalance(c *fiber.Ctx) error {
	issuer, _ := helpers.GetIssuer(c)

	var BalanceRequest struct {
		Amount float64 `json:"amount"`
	}

	if err := c.BodyParser(&BalanceRequest); err != nil {
		return err
	}

	var user model.User

	var avaliableBalance float64
	if err := database.Conn.Model(&user).Where("id = ?", issuer).Pluck("balance", &avaliableBalance).Error; err != nil {
		return err
	}

	new_balance := avaliableBalance + BalanceRequest.Amount

	if err := database.Conn.Model(&user).Where("id = ?", issuer).Update("Balance", new_balance).Error; err != nil {
		return err
	}

	CurrencyTransaction(issuer, "Deposit", avaliableBalance, BalanceRequest.Amount)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"username":    user.Username,
		"old_balance": avaliableBalance,
		"new_balance": user.Balance,
	})
}
