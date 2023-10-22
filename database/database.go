package database

import (
	"fmt"

	"doflamingo/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Conn *gorm.DB

func Connect() {
	Database()
	Migrate()
	fmt.Println("Database connection success!")
}

func Database() {
	dsn := "host=postgres user=postgres password=393406 dbname=doflamingo port=5432 sslmode=disable TimeZone=Europe/Istanbul"
	var err error
	Conn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Database error: " + err.Error())
	}
}

func Migrate() {

	Conn.AutoMigrate(
		&model.Product{},
		&model.User{},
		&model.CurrencyTransaction{},
	)

	fmt.Println("Migrate successful!")

}
