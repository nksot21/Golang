package routes

import (
	"connectdb/src/handler"

	"github.com/gofiber/fiber/v2"
)

func SetUpAuthorRoutes(app *fiber.App) {
	// POST: /signup
	app.Post("/signup", handler.SignUp)
	// POST: /signin
	app.Post("/signin", handler.SignIn)
}
