package routes

import (
	"chatdemo/src/handler"

	"github.com/gofiber/fiber/v2"
)

func SetUpUserRoutes(app *fiber.App) {
	app.Get("/user", handler.SignIn)
	app.Post("/user", handler.SignUp)
	app.Get("/home", func(c *fiber.Ctx) error {
		return c.SendString("helllo")
	})
}
