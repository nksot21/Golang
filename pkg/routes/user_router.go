package routes

import (
	"mental-health-api/handler"

	"github.com/gofiber/fiber/v2"
)

func UserRouter(a *fiber.App) {
	router := a.Group("/user")
	router.Get("/get-info", handler.GetUser)
	router.Get("/all", handler.GetUsers)
	router.Post("/login", handler.Login)
	router.Post("/", handler.CreateUser)
	router.Put("/", handler.UpdateUser)
	router.Delete("/", handler.DeleteUser)
}
