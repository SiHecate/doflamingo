package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string  `json:"username" gorm:"unique"`
	Password string  `json:"password"`
	Balance  float64 `json:"balance"`
	Type     string  `json:"type"`
}
