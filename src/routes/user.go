package routes

import (
	"connectdb/src/handler"

	"github.com/gofiber/fiber/v2"
)

func SetUpUserRoutes(user fiber.Router) {
	// GET: /user
	user.Get("/", handler.GetUsers)
	// GET: /user/:id
	user.Get("/:id", handler.GetUser)
	// PUT: /user/:id
	user.Put("/:id", handler.UpdateUser)
	// DELETE: /user/:id
	user.Delete("/:id", handler.DeleteUser)
}
