package controllers

import (
	"doflamingo/database"
	"doflamingo/model"

	"github.com/gofiber/fiber/v2"
)

func ProductAdd(c *fiber.Ctx) error {
	var data struct {
		Type        string  `json:"product_type"`
		Title       string  `json:"product_title"`
		Description string  `json:"product_description"`
		Price       float64 `json:"product_price"`
		Rate        uint8   `json:"product_rate"`
	}

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	db_response := model.Product{
		Type:        data.Type,
		Title:       data.Title,
		Description: data.Description,
		Price:       data.Price,
		Rate:        data.Rate,
	}

	if err := database.Conn.Create(&db_response).Error; err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message":       "successful!",
		"added_product": db_response,
	})
}

func ProductShowAll(c *fiber.Ctx) error {

	var products []model.Product

	database.Conn.Model(&products).Find(&products)

	return c.JSON(fiber.Map{
		"message":  "All products exist",
		"products": products,
	})
}

func ProductUpdate(c *fiber.Ctx) error {
	var data struct {
		ProductID uint   `json:"product_id"`
		Column    string `json:"column"`
		Update    string `json:"update"`
	}

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var products model.Product
	if err := database.Conn.Model(&products).Where("id = ?", data.ProductID).Update(data.Column, data.Update).Error; err != nil {
		return c.JSON(fiber.Map{
			"message": "Database update error!",
		})
	}

	return c.JSON(fiber.Map{
		"message": "successful",
		"product": products,
	})
}
