package router

import (
	"doflamingo/controllers"
	"doflamingo/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	userAuth := app.Group("/userAuth")
	userAuth.Post("/register", controllers.UserRegister)
	userAuth.Post("/login", controllers.UserLogin)

	user := app.Group("/user")
	user.Use(middlewares.JWTMiddleware())
	user.Get("/user", controllers.GetLoggedInUser)
	user.Post("/addBalance", controllers.AddBalance)

	merchantAuth := app.Group("/merchantAuth")
	merchantAuth.Post("/register", controllers.MerchantRegister)
	merchantAuth.Post("/login", controllers.MerchantLogin)

	product := app.Group("/product")
	product.Post("/add", controllers.ProductAdd)
	product.Post("/update", controllers.ProductUpdate)
	product.Get("/all", controllers.ProductShowAll)
}
