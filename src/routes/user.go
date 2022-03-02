package routes

import (
	"chatdemo/src/handler"

	"github.com/gofiber/fiber/v2"
)

func SetUpUserRoutes(app *fiber.App) {
	app.Get("/user", handler.SignIn)
	app.Post("/user", handler.SignUp)
}
