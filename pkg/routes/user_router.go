package routes

import (
	"mental-health-api/handler"

	"github.com/gofiber/fiber/v2"
)

func UserRouter(a *fiber.App) {
	router := a.Group("/user")
	router.Get("/get-info", handler.GetUser)
	router.Post("/", handler.CreateUser)
	router.Put("/", handler.UpdateUser)
	router.Delete("/", handler.DeleteUser)
}
