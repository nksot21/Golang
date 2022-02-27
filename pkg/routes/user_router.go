package routes

import (
	"mental-health-api/handler"

	"github.com/gofiber/fiber/v2"
)

func UserRouter(a *fiber.App) {
	router := a.Group("/user")
	router.Post("/", handler.CreateUser)
}
