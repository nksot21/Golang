package routes

import (
	"github.com/gofiber/fiber/v2"
	"mental-health-api/handler"
)

func UserRouter(a *fiber.App) {
	router := a.Group("/user")
	router.Get("/", handler.GetUser)
	router.Post("/", handler.CreateUser)
	router.Put("/", handler.UpdateUser)
}
