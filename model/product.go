package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Type        string  `json:"product_type"`
	Title       string  `json:"product_title" gorm:"unique"`
	Description string  `json:"product_description"`
	Price       float64 `json:"product_price"`
	Rate        uint8   `json:"product_rate"`
}
