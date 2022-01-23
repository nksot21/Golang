package routes

import (
	"connectdb/src/handler"

	"github.com/gofiber/fiber/v2"
)

func SetUpUserRoutes(app *fiber.App) {
	// POST: /newuser
	app.Post("/newuser", handler.CreateUser)
	// GET: /user
	app.Get("/user", handler.GetUsers)
	// GET: /user/:id
	app.Get("/user/:id", handler.GetUser)
	// PUT: /user/:id
	app.Put("/user/:id", handler.UpdateUser)
	// DELETE: /user/:id
	app.Delete("/user/:id", handler.DeleteUser)
}
