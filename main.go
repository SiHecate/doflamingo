package main

import (
	"doflamingo/database"
	"doflamingo/router"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	database.Connect()
	app := fiber.New()

	logFile, err := os.Create("api.log")
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)
	logFile.Sync()

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	router.Router(app)

	app.Listen(":8080")

	if err != nil {
		log.Fatal(err)
	}
}
